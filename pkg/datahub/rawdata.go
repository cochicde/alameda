package datahub

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/status"
	"prophetstor.com/api/datahub/rawdata"
)

func (p *Client) ReadRawdata(request *rawdata.ReadRawdataRequest) (*rawdata.ReadRawdataResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ReadRawdata(context.Background(), request)
}

func (p *Client) WriteRawdata(request *rawdata.WriteRawdataRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.WriteRawdata(context.Background(), request)
}
