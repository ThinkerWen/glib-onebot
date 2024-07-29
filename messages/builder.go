package messages

func NewTextMsg(text string) Message {
	data := new(TextData)
	data.Text = text
	return Message{
		Type: "text",
		Data: data,
	}
}

func NewFaceMsg(id string) Message {
	data := new(FaceData)
	data.Id = id
	return Message{
		Type: "face",
		Data: data,
	}
}

func NewImageMsg(file, fileType, url string, cache, proxy, timeout int) Message {
	data := new(ImageData)
	data.File = file
	data.Type = fileType
	data.Url = url
	data.Cache = cache
	data.Proxy = proxy
	data.Timeout = timeout
	return Message{
		Type: "image",
		Data: data,
	}
}

func NewRecordMsg(file string, magic int, url string, cache, proxy, timeout int) Message {
	data := new(RecordData)
	data.File = file
	data.Magic = magic
	data.Url = url
	data.Cache = cache
	data.Proxy = proxy
	data.Timeout = timeout
	return Message{
		Type: "record",
		Data: data,
	}
}

func NewVideoMsg(file, url string, cache, proxy, timeout int) Message {
	data := new(VideoData)
	data.File = file
	data.Url = url
	data.Cache = cache
	data.Proxy = proxy
	data.Timeout = timeout
	return Message{
		Type: "video",
		Data: data,
	}
}

func NewAtMsg(qq string) Message {
	data := new(AtData)
	data.QQ = qq
	return Message{
		Type: "at",
		Data: data,
	}
}

func NewRpsMsg() Message {
	data := new(RpsData)
	return Message{
		Type: "rps",
		Data: data,
	}
}

func NewDiceMsg() Message {
	data := new(DiceData)
	return Message{
		Type: "dice",
		Data: data,
	}
}

func NewShakeMsg() Message {
	data := new(ShakeData)
	return Message{
		Type: "shake",
		Data: data,
	}
}

func NewPokeMsg(pokeType, id, name string) Message {
	data := new(PokeData)
	data.Type = pokeType
	data.Id = id
	data.Name = name
	return Message{
		Type: "poke",
		Data: data,
	}
}

func NewAnonymousMsg(id int64, name, flag string) Message {
	data := new(AnonymousData)
	data.Id = id
	data.Name = name
	data.Flag = flag
	return Message{
		Type: "anonymous",
		Data: data,
	}
}

func NewShareMsg(url, title, content, image string) Message {
	data := new(ShareData)
	data.Url = url
	data.Title = title
	data.Content = content
	data.Image = image
	return Message{
		Type: "share",
		Data: data,
	}
}

func NewContactMsg(contactType, id string) Message {
	data := new(ContactData)
	data.Type = contactType
	data.Id = id
	return Message{
		Type: "contact",
		Data: data,
	}
}

func NewLocationMsg(lat, lon, title, content string) Message {
	data := new(LocationData)
	data.Lat = lat
	data.Lon = lon
	data.Title = title
	data.Content = content
	return Message{
		Type: "location",
		Data: data,
	}
}

func NewMusicMsg(musicType, id, url, audio, title, content, image string) Message {
	data := new(MusicData)
	data.Type = musicType
	data.Id = id
	data.Url = url
	data.Audio = audio
	data.Title = title
	data.Content = content
	data.Image = image
	return Message{
		Type: "music",
		Data: data,
	}
}

func NewReplyMsg(id string) Message {
	data := new(ReplyData)
	data.Id = id
	return Message{
		Type: "reply",
		Data: data,
	}
}

func NewForwardMsg(id string) Message {
	data := new(ForwardData)
	data.Id = id
	return Message{
		Type: "forward",
		Data: data,
	}
}

func NewNodeMsg(id, userId, nickname string, content Message) Message {
	data := new(NodeData)
	data.Id = id
	data.UserId = userId
	data.Nickname = nickname
	data.Content = content
	return Message{
		Type: "node",
		Data: data,
	}
}

func NewXmlMsg(dataString string) Message {
	data := new(XmlData)
	data.Data = dataString
	return Message{
		Type: "xml",
		Data: data,
	}
}

func NewJsonMsg(dataString string) Message {
	data := new(JsonData)
	data.Data = dataString
	return Message{
		Type: "json",
		Data: data,
	}
}
