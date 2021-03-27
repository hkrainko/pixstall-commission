package grpc

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"io"
	"pixstall-commission/app/image/grpc/proto"
	model2 "pixstall-commission/domain/file/model"
	"pixstall-commission/domain/image"
)

type grpcImageRepository struct {
	grpcConn *grpc.ClientConn
}

func NewGRPCImageRepository(grpcConn *grpc.ClientConn) image.Repo {
	return &grpcImageRepository{
		grpcConn: grpcConn,
	}
}

func (g grpcImageRepository) SaveFile(ctx context.Context, file model2.File, fileType model2.FileType, ownerID string, acl []string) (*string, error) {
	client := proto.NewFileServiceClient(g.grpcConn)

	stream, err := client.SaveFile(ctx)
	if err != nil {
		return nil, err
	}
	gFileType, err := g.gRPCFileTypeFormDomain(fileType)
	if err != nil {
		return nil, err
	}
	req := &proto.SaveFileRequest{
		Data: &proto.SaveFileRequest_MetaData{
			MetaData: &proto.MetaData{
				FileType: gFileType,
				Name:     file.Name,
				Owner: ownerID,
				Acl: acl,
			},
		},
	}
	err = stream.SendMsg(req)
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, 1024)

	for {
		n, err := file.File.Read(buffer)
		if err == io.EOF {
			break
		}
		req := &proto.SaveFileRequest{
			Data: &proto.SaveFileRequest_File{
				File: buffer[:n],
			},
		}
		err = stream.SendMsg(req)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}
	return &resp.Path, nil
}

func (g grpcImageRepository) SaveFiles(ctx context.Context, files []model2.File, fileType model2.FileType, ownerID string, acl []string) ([]string, error) {
	panic("implement me")
}

func (g grpcImageRepository) gRPCFileTypeFormDomain(dFileType model2.FileType) (proto.MetaData_FileType, error) {
	switch dFileType {
	case model2.FileTypeMessage:
		return proto.MetaData_Message, nil
	case model2.FileTypeCompletion:
		return proto.MetaData_Completion, nil
	case model2.FileTypeCommissionRef:
		return proto.MetaData_CommissionRef, nil
	case model2.FileTypeCommissionProofCopy:
		return proto.MetaData_CommissionProofCopy, nil
	case model2.FileTypeArtwork:
		return proto.MetaData_Artwork, nil
	case model2.FileTypeRoof:
		return proto.MetaData_Roof, nil
	case model2.FileTypeOpenCommission:
		return proto.MetaData_OpenCommission, nil
	case model2.FileTypeProfile:
		return proto.MetaData_Profile, nil
	default:
		return -1, errors.New("not found")
	}
}