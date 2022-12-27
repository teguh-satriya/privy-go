package server

import (
	"context"
	"fmt"

	cakesv1 "github.com/teguh-satriya/privy-go/proto/cakes/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *CakesServer) ListCakes(ctx context.Context, req *cakesv1.ListCakesRequest) (*cakesv1.ListCakesResponse, error) {
	result, err := s.listCakesService.Call(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("result: %v\n", result)

	res := &cakesv1.ListCakesResponse{
		Data: []*cakesv1.Cake{},
	}

	for _, item := range result.Cakes {
		row := &cakesv1.Cake{
			Id:          int32(item.ID),
			Title:       item.Title,
			Description: item.Description,
			Rating:      int32(item.Rating),
			Image:       item.Image,
			CreatedAt:   timestamppb.New(item.CreatedAt),
			UpdatedAt:   timestamppb.New(item.UpdatedAt),
		}

		res.Data = append(res.Data, row)
	}

	return res, nil
}
