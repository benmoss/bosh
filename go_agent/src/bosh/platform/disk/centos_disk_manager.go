package disk

import (
	boshlog "bosh/logger"
	boshsys "bosh/system"
)

type centosDiskManager struct {
	partitioner Partitioner
	formatter   Formatter
	mounter     Mounter
	finder      Finder
}

func NewCentosDiskManager(logger boshlog.Logger, runner boshsys.CmdRunner, fs boshsys.FileSystem, finder Finder) (manager Manager) {
	return centosDiskManager{
		partitioner: newSfdiskPartitioner(logger, runner),
		formatter:   newLinuxFormatter(runner, fs),
		mounter:     newLinuxMounter(runner, fs),
		finder:      finder,
	}
}

func (m centosDiskManager) GetPartitioner() Partitioner {
	return m.partitioner
}

func (m centosDiskManager) GetFormatter() Formatter {
	return m.formatter
}

func (m centosDiskManager) GetMounter() Mounter {
	return m.mounter
}

func (m centosDiskManager) GetFinder() Finder {
	return m.finder
}
