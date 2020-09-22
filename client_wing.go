package torrent

import (
	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage"
	"os"
	"path/filepath"
)

func WingFakeClient() *Client {
	cnf := NewDefaultClientConfig()
	tempDir := filepath.Join(os.TempDir(), "fake_torrent_client")
	os.MkdirAll(tempDir, 0777)
	cnf.DefaultStorage = storage.NewBoltDB(tempDir)
	fake := &Client{
		config:         cnf,
		defaultStorage: storage.NewClient(cnf.DefaultStorage),
		torrents:       make(map[metainfo.Hash]*Torrent),
		logger:         log.Discard,
	}
	return fake
}

// File的扩展
func (f *File) FirstPieceIndex() int {
	return f.firstPieceIndex()
}

func (f *File) EndPieceIndex() int {
	return f.endPieceIndex()
}
