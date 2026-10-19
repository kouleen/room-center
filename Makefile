thrift=thrift
idl_path = idl/room-center

gen_service:
	kitex -module github.com/kouleen/room-center -service room.service  $(idl_path)/room.thrift
