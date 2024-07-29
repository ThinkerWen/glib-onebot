package messages

import "encoding/json"

type MsgFactory struct {
	data []byte
}

func MsgParser(msg Message) *MsgFactory {
	if data, err := json.Marshal(msg); err != nil {
		return nil
	} else {
		factory := new(MsgFactory)
		factory.data = data
		return factory
	}
}

func (m MsgFactory) AsText() TextData {
	var data TextData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return TextData{}
	}
	return data
}

func (m MsgFactory) AsFace() FaceData {
	var data FaceData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return FaceData{}
	}
	return data
}

func (m MsgFactory) AsImage() ImageData {
	var data ImageData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return ImageData{}
	}
	return data
}

func (m MsgFactory) AsRecord() RecordData {
	var data RecordData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return RecordData{}
	}
	return data
}

func (m MsgFactory) AsVideo() VideoData {
	var data VideoData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return VideoData{}
	}
	return data
}

func (m MsgFactory) AsAt() AtData {
	var data AtData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return AtData{}
	}
	return data
}

func (m MsgFactory) AsRps() RpsData {
	var data RpsData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return RpsData{}
	}
	return data
}

func (m MsgFactory) AsDice() DiceData {
	var data DiceData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return DiceData{}
	}
	return data
}

func (m MsgFactory) AsShake() ShakeData {
	var data ShakeData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return ShakeData{}
	}
	return data
}

func (m MsgFactory) AsPoke() PokeData {
	var data PokeData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return PokeData{}
	}
	return data
}

func (m MsgFactory) AsAnonymous() AnonymousData {
	var data AnonymousData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return AnonymousData{}
	}
	return data
}

func (m MsgFactory) AsShare() ShareData {
	var data ShareData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return ShareData{}
	}
	return data
}

func (m MsgFactory) AsContact() ContactData {
	var data ContactData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return ContactData{}
	}
	return data
}

func (m MsgFactory) AsLocation() LocationData {
	var data LocationData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return LocationData{}
	}
	return data
}

func (m MsgFactory) AsMusic() MusicData {
	var data MusicData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return MusicData{}
	}
	return data
}

func (m MsgFactory) AsReply() ReplyData {
	var data ReplyData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return ReplyData{}
	}
	return data
}

func (m MsgFactory) AsForward() ForwardData {
	var data ForwardData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return ForwardData{}
	}
	return data
}

func (m MsgFactory) AsNode() NodeData {
	var data NodeData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return NodeData{}
	}
	return data
}

func (m MsgFactory) AsXml() XmlData {
	var data XmlData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return XmlData{}
	}
	return data
}

func (m MsgFactory) AsJson() JsonData {
	var data JsonData

	if err := json.Unmarshal(m.data, &data); err != nil {
		return JsonData{}
	}
	return data
}
