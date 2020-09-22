package torrent

import "github.com/anacrolix/torrent/metainfo"

//wing的扩展

func (t *Torrent) WingMergeFromMetainfo(mi *metainfo.MetaInfo) {
	t.metainfo = *mi
	//t.info = mi

}
