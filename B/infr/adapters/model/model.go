package model

type User struct {
	Id      string `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Phone   string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
}
type UserUpdate struct {
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Phone   string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
}
