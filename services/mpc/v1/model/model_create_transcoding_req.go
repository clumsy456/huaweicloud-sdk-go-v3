package model

import (
	"encoding/json"

	"strings"
)

type CreateTranscodingReq struct {
	Input *ObsObjInfo `json:"input,omitempty"`

	Output *ObsObjInfo `json:"output"`
	// 转码模板ID，没带av_parameter参数时，必须带该参数，数组，每一路转码输出对应一个转码配置模板ID，最多支持9个模板ID。  多个转码模板中如下参数可变，其他都必须一致：  - 视频bitrate，height，width。

	TransTemplateId *[]int32 `json:"trans_template_id,omitempty"`
	// 转码参数。  若同时设置“trans_template_id”和此参数，则优先使用此参数进行转码，不带trans_template_id时，该参数必选。

	AvParameters *[]AvParameters `json:"av_parameters,omitempty"`
	// 输出文件名称，每一路转码输出对应一个名称，需要与转码模板ID数组的顺序对应。  - 若设置该参数，表示输出文件按该参数命名。 - 若不设置该参数，表示输出文件按默认方式命名。

	OutputFilenames *[]string `json:"output_filenames,omitempty"`
	// 用户自定义数据，该字段可在查询接口或消息通知中按原内容透传给用户。

	UserData *string `json:"user_data,omitempty"`
	// 图片水印参数，数组，最多支持20个成员。

	Watermarks *[]WatermarkRequest `json:"watermarks,omitempty"`

	Thumbnail *Thumbnail `json:"thumbnail,omitempty"`

	DigitalWatermark *DigitalWatermark `json:"digital_watermark,omitempty"`
	// 项目ID

	ProjectId *string `json:"project_id,omitempty"`
	// 是否vip用户

	VipUser *bool `json:"vip_user,omitempty"`
	// 任务Id

	TaskId *string `json:"task_id,omitempty"`
	// 租户域名

	DomainName *string `json:"domain_name,omitempty"`
	// 租户Id

	TenantProjectId *string `json:"tenant_project_id,omitempty"`
	// 任务优先级，取值如下： - 9代表高优先级。 - 6代表中优先级，默认为6。  暂时只支持6和9。

	Priority *int32 `json:"priority,omitempty"`

	Audit *Audit `json:"audit,omitempty"`

	Subtitle *Subtitle `json:"subtitle,omitempty"`

	SpecialEffect *SpecialEffect `json:"special_effect,omitempty"`

	Encryption *Encryption `json:"encryption,omitempty"`

	Crop *Crop `json:"crop,omitempty"`

	AudioTrack *AudioTrack `json:"audio_track,omitempty"`

	MultiAudio *MultiAudio `json:"multi_audio,omitempty"`

	VideoProcess *VideoProcess `json:"video_process,omitempty"`

	AudioProcess *AudioProcess `json:"audio_process,omitempty"`

	QualityEnhance *QualityEnhance `json:"quality_enhance,omitempty"`

	SystemProcess *SystemProcess `json:"system_process,omitempty"`

	TemplateExtend *TemplateExtend `json:"template_extend,omitempty"`
}

func (o CreateTranscodingReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTranscodingReq struct{}"
	}

	return strings.Join([]string{"CreateTranscodingReq", string(data)}, " ")
}
