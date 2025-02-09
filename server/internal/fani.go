package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"fani/api"
	faniv1 "fani/gen/fani/v1"

	"connectrpc.com/connect"
)

type FaniServer struct {
	api.Settings
}

func (s *FaniServer) Start(ctx context.Context) {
	for _, pluginPath := range s.Settings.PluginPaths {
		_ = pluginPath
	}
}

func (s *FaniServer) ListStaging(
	ctx context.Context,
	req *connect.Request[faniv1.ListStagingRequest],
) (*connect.Response[faniv1.ListStagingResponse], error) {
	log.Printf("%s %+v", getCurrentFuncName(), req.Msg)
	var stagingFiles []*faniv1.FileObject
	files, err := os.ReadDir(s.StagingDirectory)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.StagingDirectory, file.Name()))
		if err != nil {
			return nil, err
		}
		stagingFiles = append(stagingFiles, &faniv1.FileObject{
			Name:      file.Name(),
			Thumbnail: data,
			Attributes: map[string]string{
				"foo": "bar",
			},
		})
	}
	return connect.NewResponse(&faniv1.ListStagingResponse{
		Files: stagingFiles,
	}), nil
}

func (s *FaniServer) ListAlbums(
	ctx context.Context,
	req *connect.Request[faniv1.ListAlbumsRequest],
) (*connect.Response[faniv1.ListAlbumsResponse], error) {
	log.Printf("%s %+v", getCurrentFuncName(), req.Msg)

	// gather streams for images from the following IDs

	var responseFiles []*faniv1.FileObject
	for range 50 {
		_ = req.Msg.Page
		responseFiles = append(responseFiles, &faniv1.FileObject{
			Attributes: map[string]string{
				"awake": "test!",
			},
		})
	}
	return connect.NewResponse(&faniv1.ListAlbumsResponse{
		Files: responseFiles,
	}), nil
}

func (s *FaniServer) MoveFiles(
	ctx context.Context,
	req *connect.Request[faniv1.MoveFilesRequest],
) (*connect.Response[faniv1.MoveFilesResponse], error) {
	log.Printf("%s %+v", getCurrentFuncName(), req.Msg)
	for range req.Msg.Items {
		// TODO:
	}
	return connect.NewResponse(&faniv1.MoveFilesResponse{
		Status: 200,
	}), nil
}

func (s *FaniServer) GetMedia(
	ctx context.Context,
	req *connect.Request[faniv1.GetMediaRequest],
) (*connect.Response[faniv1.GetMediaResponse], error) {
	log.Printf("%s %+v", getCurrentFuncName(), req.Msg)
	return connect.NewResponse(&faniv1.GetMediaResponse{}), nil
}

func getCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}
