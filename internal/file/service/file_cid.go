package service

import (
	"bytes"
	"context"
	"ipfs-file-api/internal/file/entity"
	"ipfs-file-api/pkg/ipfs"
	"ipfs-file-api/pkg/tools"
)

type fileCIDService struct {
	fileCIDRep entity.FileCIDRepository
}

func NewFileCIDService(fileCIDRep entity.FileCIDRepository) entity.FileCIDService {
	return &fileCIDService{
		fileCIDRep: fileCIDRep,
	}
}

func (s *fileCIDService) GetFileCID(ctx context.Context, id uint) (*[]byte, error) {
	fileCID, err := s.fileCIDRep.FindFirst(ctx, "id", id)
	if err != nil {
		return nil, err
	}
	buf, err := ipfs.DownloadBytes(*fileCID.CID)
	if err != nil {
		return nil, err
	}

	return tools.GetPointer(buf), nil
}

func (s *fileCIDService) CreateFileCID(ctx context.Context, buf *bytes.Buffer, name string) (*entity.FileCID, error) {
	cid, err := ipfs.UploadBytes(buf)
	if err != nil {
		return nil, err
	}

	return s.fileCIDRep.Create(ctx, entity.FileCID{Name: name, CID: tools.GetPointer(cid)})
}
