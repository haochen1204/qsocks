package proxy

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/net-byte/qsocks/common/cipher"
	"github.com/net-byte/qsocks/config"
	"github.com/quic-go/quic-go"
)

var _tlsConf *tls.Config
var _lock sync.Mutex

func ConnectServer(config config.Config) (quic.Connection, error) {
	_lock.Lock()
	if _tlsConf == nil {
		var err error
		_tlsConf, err = config.GetClientTLSConfig()
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	_lock.Unlock()
	quicConfig := &quic.Config{
		HandshakeIdleTimeout: time.Second * 10,
		MaxIdleTimeout:       time.Second * 30,
	}
	session, err := quic.DialAddr(context.Background(), config.ServerAddr, _tlsConf, quicConfig)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return session, nil
}

func Handshake(network string, host string, port string, session quic.Connection) bool {
	// handshake
	req := &RequestAddr{}
	req.Network = network
	req.Host = host
	req.Port = port
	req.Timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	req.Random = cipher.Random()
	data, err := req.MarshalBinary()
	if err != nil {
		log.Printf("[client] failed to encode request addr %v", err)
		return false
	}
	stream, err := session.OpenUniStreamSync(context.Background())
	if err != nil {
		log.Println(err)
		return false
	}
	stream.Write(data)
	return true
}

func Copy(destination io.WriteCloser, source io.ReadCloser) {
	if destination == nil || source == nil {
		return
	}
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}
