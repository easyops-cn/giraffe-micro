package easyopsrest

import (
	"github.com/easyops-cn/giraffe-micro"
	"reflect"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
)

type Instance struct {
	ObjectId             string        `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Data                 *types.Struct `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}

type InstanceID struct {
	ObjectId             string   `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	InstanceId           string   `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Version              int32    `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstanceID) Reset()         { *m = InstanceID{} }
func (m *InstanceID) String() string { return proto.CompactTextString(m) }
func (*InstanceID) ProtoMessage()    {}

func Test_request_getVerbPath(t *testing.T) {
	type canNotInterface struct {
		canNotInterface string
	}

	type notSupportParameter struct {
		NotSupportParameter uint64
	}

	type fields struct {
		message  interface{}
		propMap  map[string]*proto.Properties
		contract giraffe.Contract
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		want1   string
		wantErr bool
	}{
		{
			name: "Test_HappyPath",
			fields: fields{
				message: &InstanceID{
					ObjectId:   "HOST",
					InstanceId: "5c6f6cf0d8079",
					Version:    1,
				},
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"url_pattern": "GET /object/:object_id/instance/:instance_id/:version",
					},
				},
			},
			want:    "GET",
			want1:   "/object/HOST/instance/5c6f6cf0d8079/1",
			wantErr: false,
		},
		{
			name: "Test_InvalidParameter",
			fields: fields{
				message: &InstanceID{
					ObjectId:   "HOST",
					InstanceId: "5c6f6cf0d8079",
					Version:    1,
				},
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"url_pattern": "GET /object/:object_id/instance/:instance_id/:version_invalid",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test_CanNotInterface",
			fields: fields{
				message: &canNotInterface{
					canNotInterface: "abc",
				},
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"url_pattern": "GET /error/:canNotInterface",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test_URLPatternNotSet",
			fields: fields{
				message: &InstanceID{
					ObjectId:   "HOST",
					InstanceId: "5c6f6cf0d8079",
					Version:    1,
				},
				contract: &giraffe.MethodDesc{
					ServiceName: "logic.cmdb",
					MethodName:  "GetInstance",
					MetaData:    map[string]interface{}{},
				},
			},
			want:    "GET",
			want1:   "/logic.cmdb/GetInstance",
			wantErr: false,
		},
		{
			name: "Test_URLPatternWasEmptyString",
			fields: fields{
				message: &InstanceID{
					ObjectId:   "HOST",
					InstanceId: "5c6f6cf0d8079",
					Version:    1,
				},
				contract: &giraffe.MethodDesc{
					ServiceName: "logic.cmdb",
					MethodName:  "GetInstance",
					MetaData: map[string]interface{}{
						"url_pattern": "",
					},
				},
			},
			want:    "GET",
			want1:   "/logic.cmdb/GetInstance",
			wantErr: false,
		},
		{
			name: "Test_InvalidURLPatternFormat",
			fields: fields{
				message: &InstanceID{
					ObjectId:   "HOST",
					InstanceId: "5c6f6cf0d8079",
					Version:    1,
				},
				contract: &giraffe.MethodDesc{
					ServiceName: "logic.cmdb",
					MethodName:  "GetInstance",
					MetaData: map[string]interface{}{
						"url_pattern": "/",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test_NotSupportParameterType",
			fields: fields{
				message: &notSupportParameter{
					NotSupportParameter: 123,
				},
				contract: &giraffe.MethodDesc{
					ServiceName: "logic.cmdb",
					MethodName:  "GetInstance",
					MetaData: map[string]interface{}{
						"url_pattern": "GET /:NotSupportParameter",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &request{
				message:  tt.fields.message,
				propMap:  tt.fields.propMap,
				contract: tt.fields.contract,
			}
			got, got1, err := r.getVerbPath()
			if (err != nil) != tt.wantErr {
				t.Errorf("request.getVerbPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("request.getVerbPath() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("request.getVerbPath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_request_getFieldByOrigName(t *testing.T) {
	type CanNotInterface struct {
		unexportedField string
	}

	type Nilable struct {
		NilableField interface{}
	}

	type fields struct {
		message  interface{}
		propMap  map[string]*proto.Properties
		contract giraffe.Contract
	}
	type args struct {
		origName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test_HappyPath",
			fields: fields{
				message: &InstanceID{
					ObjectId:   "HOST",
					InstanceId: "5c6f6cf0d8079",
				},
			},
			args: args{
				origName: "object_id",
			},
			want:    "HOST",
			wantErr: false,
		},
		{
			name: "Test_NoMatchFieldByOrigName",
			fields: fields{
				message: &InstanceID{
					ObjectId:   "HOST",
					InstanceId: "5c6f6cf0d8079",
				},
			},
			args: args{
				origName: "ObjectId",
			},
			wantErr: true,
		},
		{
			name: "Test_NoMatchFieldByName",
			fields: fields{
				message: &InstanceID{
					ObjectId:   "HOST",
					InstanceId: "5c6f6cf0d8079",
				},
				propMap: map[string]*proto.Properties{
					"ObjectId": {
						Name: "object_id",
					},
				},
			},
			args: args{
				origName: "ObjectId",
			},
			wantErr: true,
		},
		{
			name: "Test_CanNotInterface",
			fields: fields{
				message: &CanNotInterface{
					unexportedField: "unexported field",
				},
			},
			args: args{
				origName: "unexportedField",
			},
			wantErr: true,
		},
		{
			name: "Test_NilField",
			fields: fields{
				message: &Nilable{
					NilableField: nil,
				},
			},
			args: args{
				origName: "NilableField",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Test_MessageNotStruct",
			fields: fields{
				message: nil,
			},
			args: args{
				origName: "any",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &request{
				message:  tt.fields.message,
				propMap:  tt.fields.propMap,
				contract: tt.fields.contract,
			}
			got, err := r.getFieldByOrigName(tt.args.origName)
			if (err != nil) != tt.wantErr {
				t.Errorf("request.getFieldByOrigName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("request.getFieldByOrigName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_request_getDataField(t *testing.T) {
	type NotProtoMessage struct {
		name string
	}
	type Wrapper struct {
		Instance        Instance
		NotProtoMessage NotProtoMessage
	}

	type fields struct {
		message  interface{}
		propMap  map[string]*proto.Properties
		contract giraffe.Contract
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Test_HappyPath",
			fields: fields{
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"data_field": "data",
					},
				},
				message: &Instance{
					ObjectId: "HOST",
					Data: &types.Struct{
						Fields: map[string]*types.Value{
							"instance_id": {Kind: &types.Value_StringValue{StringValue: "5c6f6cf0d8079"}},
							"name":        {Kind: &types.Value_StringValue{StringValue: "123"}},
						},
					},
				},
			},
			want:    []byte("{\"instance_id\":\"5c6f6cf0d8079\",\"name\":\"123\"}"),
			wantErr: false,
		},
		{
			name: "Test_DataFieldWasEmptyString",
			fields: fields{
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"data_field": "",
					},
				},
				message: &Instance{
					ObjectId: "HOST",
					Data: &types.Struct{
						Fields: map[string]*types.Value{
							"instance_id": {Kind: &types.Value_StringValue{StringValue: "5c6f6cf0d8079"}},
							"name":        {Kind: &types.Value_StringValue{StringValue: "123"}},
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Test_NoDataField",
			fields: fields{
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{},
				},
				message: &Instance{
					ObjectId: "HOST",
					Data: &types.Struct{
						Fields: map[string]*types.Value{
							"instance_id": {Kind: &types.Value_StringValue{StringValue: "5c6f6cf0d8079"}},
							"name":        {Kind: &types.Value_StringValue{StringValue: "123"}},
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Test_WrongDataField",
			fields: fields{
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"data_field": "Data",
					},
				},
				message: &Instance{
					ObjectId: "HOST",
					Data: &types.Struct{
						Fields: map[string]*types.Value{
							"instance_id": {Kind: &types.Value_StringValue{StringValue: "5c6f6cf0d8079"}},
							"name":        {Kind: &types.Value_StringValue{StringValue: "123"}},
						},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test_DataFieldWasStruct",
			fields: fields{
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"data_field": "Instance",
					},
				},
				message: &Wrapper{
					Instance: Instance{
						ObjectId: "HOST",
						Data: &types.Struct{
							Fields: map[string]*types.Value{
								"instance_id": {Kind: &types.Value_StringValue{StringValue: "5c6f6cf0d8079"}},
								"name":        {Kind: &types.Value_StringValue{StringValue: "123"}},
							},
						},
					},
				},
			},
			want:    []byte("{\"object_id\":\"HOST\",\"data\":{\"instance_id\":\"5c6f6cf0d8079\",\"name\":\"123\"}}"),
			wantErr: false,
		},
		{
			name: "Test_DataFieldWasNotStructOrPointerOfStruct",
			fields: fields{
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"data_field": "object_id",
					},
				},
				message: &Instance{
					ObjectId: "HOST",
					Data: &types.Struct{
						Fields: map[string]*types.Value{
							"instance_id": {Kind: &types.Value_StringValue{StringValue: "5c6f6cf0d8079"}},
							"name":        {Kind: &types.Value_StringValue{StringValue: "123"}},
						},
					},
				},
			},
			want:    []byte("\"HOST\""),
			wantErr: false,
		},
		{
			name: "Test_DataFieldWasNotProtoMessage",
			fields: fields{
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"data_field": "NotProtoMessage",
					},
				},
				message: &Wrapper{
					NotProtoMessage: NotProtoMessage{
						name: "123",
					},
				},
			},
			want:    []byte("{}"),
			wantErr: false,
		},
		{
			name: "Test_DataFieldWasNil",
			fields: fields{
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"data_field": "data",
					},
				},
				message: &Instance{
					ObjectId: "HOST",
					Data:     nil,
				},
			},
			wantErr: true,
		},
		{
			name: "Test_DataFieldWasNilProtoMessage",
			fields: fields{
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"data_field": "data",
					},
				},
				message: &Instance{
					ObjectId: "HOST",
					Data:     (*types.Struct)(nil),
				},
			},
			wantErr: true,
		},
		{
			name: "Test_Empty",
			fields: fields{
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"data_field": "Data",
					},
				},
				message: &types.Empty{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &request{
				message:  tt.fields.message,
				propMap:  tt.fields.propMap,
				contract: tt.fields.contract,
			}
			got, err := r.getDataField()
			if (err != nil) != tt.wantErr {
				t.Errorf("request.getDataField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("request.getDataField() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func TestNewRequest(t *testing.T) {
	client := rest.NewRESTClient(nil)

	type NotProtoMessage struct {
		ObjectId string                 `json:"object_id,omitempty"`
		Data     map[string]interface{} `json:"data,omitempty"`
	}
	type args struct {
		c        rest.Interface
		contract giraffe.Contract
		in       interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *rest.Request
		wantErr bool
	}{
		{
			name: "Test_HappyPath",
			args: args{
				c: client,
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"url_pattern": "POST /object/:object_id/instance",
						"data_field":  "data",
					},
				},
				in: &Instance{
					ObjectId: "HOST",
					Data: &types.Struct{
						Fields: map[string]*types.Value{
							"name": {Kind: &types.Value_StringValue{StringValue: "123"}},
						},
					},
				},
			},
			want: client.Verb("POST", "/object/HOST/instance").
				Header("Content-Type", "application/json").
				Body([]byte("{\"name\":\"123\"}")),
			wantErr: false,
		},
		{
			name: "Test_HappyPath2",
			args: args{
				c: client,
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"url_pattern": "POST /instance",
						"data_field":  "",
					},
				},
				in: &Instance{
					ObjectId: "HOST",
					Data: &types.Struct{
						Fields: map[string]*types.Value{
							"name": {Kind: &types.Value_StringValue{StringValue: "123"}},
						},
					},
				},
			},
			want: client.Verb("POST", "/instance").
				Header("Content-Type", "application/json").
				Body([]byte("{\"object_id\":\"HOST\",\"data\":{\"name\":\"123\"}}")),
			wantErr: false,
		},
		{
			name: "Test_HappyPath3",
			args: args{
				c: client,
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"url_pattern": "GET /object/:object_id/instance/:instance_id",
						"data_field":  "",
					},
				},
				in: &InstanceID{
					ObjectId:   "HOST",
					InstanceId: "5c6f6cf0d8079",
				},
			},
			want: client.Verb("GET", "/object/HOST/instance/5c6f6cf0d8079").
				Param("object_id", "HOST").Param("instance_id", "5c6f6cf0d8079"),
			wantErr: false,
		},
		{
			name: "Test_BodyWasNotProtoMessage",
			args: args{
				c: client,
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"url_pattern": "POST /object/:ObjectId/instance",
						"data_field":  "Data",
					},
				},
				in: &NotProtoMessage{
					ObjectId: "HOST",
					Data: map[string]interface{}{
						"name": "123",
					},
				},
			},
			want: client.Verb("POST", "/object/HOST/instance").
				Header("Content-Type", "application/json").
				Body([]byte("{\"name\":\"123\"}")),
			wantErr: false,
		},
		{
			name: "Test_WithWrongURLPattern",
			args: args{
				c: client,
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"url_pattern": "/object/:ObjectId/instance",
						"data_field":  "Data",
					},
				},
				in: &NotProtoMessage{
					ObjectId: "HOST",
					Data: map[string]interface{}{
						"name": "123",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test_WithWrongDataField",
			args: args{
				c: client,
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"url_pattern": "POST /object/:ObjectId/instance",
						"data_field":  "data",
					},
				},
				in: &NotProtoMessage{
					ObjectId: "HOST",
					Data: map[string]interface{}{
						"name": "123",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test_MessageWasNilProtoMessage",
			args: args{
				c: client,
				contract: &giraffe.MethodDesc{
					MetaData: map[string]interface{}{
						"url_pattern": "POST /",
					},
				},
				in: (*Instance)(nil),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRequest(tt.args.c, tt.args.contract, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
