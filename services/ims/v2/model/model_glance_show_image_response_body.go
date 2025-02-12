package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 镜像信息响应体
type GlanceShowImageResponseBody struct {
	// 备份ID。如果是备份创建的镜像，则填写为备份的ID，否则为空。

	BackupId string `json:"__backup_id"`
	// 镜像来源。公共镜像为空。

	DataOrigin string `json:"__data_origin"`
	// 镜像描述信息。

	Description string `json:"__description"`
	// 镜像文件的大小，单位为字节。目前取值为大于0的字符串。

	ImageSize string `json:"__image_size"`
	// 镜像后端存储类型，目前只支持uds

	ImageSourceType GlanceShowImageResponseBodyImageSourceType `json:"__image_source_type"`
	// 镜像类型，目前支持以下类型：公共镜像：gold私有镜像：private共享镜像：shared

	Imagetype GlanceShowImageResponseBodyImagetype `json:"__imagetype"`
	// 是否是注册过的镜像，取值为“true”或者“false”。

	Isregistered GlanceShowImageResponseBodyIsregistered `json:"__isregistered"`
	// 父镜像ID。公共镜像或通过文件创建的私有镜像，取值为空。

	Originalimagename string `json:"__originalimagename"`
	// 操作系统位数，一般取值为“32”或者“64”。

	OsBit GlanceShowImageResponseBodyOsBit `json:"__os_bit"`
	// 操作系统类型，目前取值Linux， Windows，Other。

	OsType GlanceShowImageResponseBodyOsType `json:"__os_type"`
	// 操作系统具体版本。

	OsVersion string `json:"__os_version"`
	// 镜像平台分类，取值为Windows，Ubuntu，RedHat，SUSE，CentOS，Debian，OpenSUSE, Oracle Linux，Fedora，Other，CoreOS和EulerOS。

	Platform GlanceShowImageResponseBodyPlatform `json:"__platform"`
	// 市场镜像的产品ID。

	Productcode string `json:"__productcode"`
	// 表示该镜像支持密集存储。如果镜像支持密集存储性能，则值为true，否则无需增加该属性。

	SupportDiskintensive string `json:"__support_diskintensive"`
	// 表示该镜像支持高计算性能。如果镜像支持高计算性能，则值为true，否则无需增加该属性。

	SupportHighperformance string `json:"__support_highperformance"`
	// 如果镜像支持KVM，取值为true，否则无需增加该属性。

	SupportKvm string `json:"__support_kvm"`
	// 表示该镜像是支持KVM虚拟化平台下的GPU类型，如果不支持KVM虚拟机下GPU类型，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存。

	SupportKvmGpuType string `json:"__support_kvm_gpu_type"`
	// 如果镜像支持KVM虚拟化下Infiniband网卡类型，取值为true。否则，无需添加该属性。该属性与“__support_xen”属性不共存。

	SupportKvmInfiniband string `json:"__support_kvm_infiniband"`
	// 表示该镜像支持超大内存。如果镜像支持超大内存，取值为true，否则无需增加该属性

	SupportLargememory string `json:"__support_largememory"`
	// 如果镜像支持XEN，取值为true，否则无需增加该属性。

	SupportXen string `json:"__support_xen"`
	// 表示该镜像是支持XEN虚拟化平台下的GPU优化类型，取值参考8.10-表 镜像支持的GPU类型说明。镜像操作系统类型请参考8.10-表 镜像类型。如果不支持XEN虚拟化下GPU类型，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存。

	SupportXenGpuType string `json:"__support_xen_gpu_type"`
	// 如果镜像支持XEN虚拟化下HANA类型，取值为true。否则，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存。

	SupportXenHana string `json:"__support_xen_hana"`
	// 目前暂时不使用。

	Checksum string `json:"checksum"`
	// 容器类型。

	ContainerFormat string `json:"container_format"`
	// 创建时间。格式为UTC时间。

	CreatedAt string `json:"created_at"`
	// 是否是删除的镜像，取值为true或者false。

	Deleted bool `json:"deleted"`
	// 删除时间。格式为UTC时间

	DeletedAt string `json:"deleted_at"`
	// 镜像的格式，目前支持vhd，zvhd、raw，qcow2,zvhd2。默认值是vhd。

	DiskFormat GlanceShowImageResponseBodyDiskFormat `json:"disk_format"`
	// 镜像文件下载和上传链接。

	File string `json:"file"`
	// 镜像ID。

	Id string `json:"id"`
	// 镜像运行需要的最小磁盘容量，单位为GB

	MinDisk int32 `json:"min_disk"`
	// 镜像运行最小内存，单位为MB。

	MinRam int32 `json:"min_ram"`
	// 镜像名称。

	Name string `json:"name"`
	// 镜像属于哪个租户。

	Owner string `json:"owner"`
	// 是否是受保护的，受保护的镜像不允许删除。取值为true或false。

	Protected bool `json:"protected"`
	// 镜像视图。

	Schema string `json:"schema"`
	// 镜像链接信息。

	Self string `json:"self"`
	// 目前暂时不使用。

	Size int64 `json:"size"`
	// 镜像状态。取值如下：queued：表示镜像元数据已经创建成功，等待上传镜像文件。saving：表示镜像正在上传文件到后端存储。deleted：表示镜像已经删除。killed：表示镜像上传错误。active：表示镜像可以正常使用。

	Status GlanceShowImageResponseBodyStatus `json:"status"`
	// 镜像标签列表，提供用户可以自定义管理私有镜像的能力。用户可以通过镜像标签接口为每个镜像增加不同的标签，在查询接口中可以根据标签进行过滤。

	Tags []string `json:"tags"`
	// 更新时间。格式为UTC时间。

	UpdatedAt string `json:"updated_at"`
	// 镜像使用环境类型：FusionCompute，Ironic，DataImage。

	VirtualEnvType GlanceShowImageResponseBodyVirtualEnvType `json:"virtual_env_type"`
	// 目前暂时不使用。

	VirtualSize int32 `json:"virtual_size"`
	// 是否被其他租户可见，取值如下：private：私有镜像public：公共镜像shared：共享镜像

	Visibility GlanceShowImageResponseBodyVisibility `json:"visibility"`
	// 表示当前镜像支持CloudInit密码/密钥注入方式，建议设置为\"true\"或者\"false\"。如果取值为\"true\"，表示该镜像不支持CloudInit注入密码/密钥，其他取值时表示支持CloudInit注入密钥/密码。

	SupportFcInject GlanceShowImageResponseBodySupportFcInject `json:"__support_fc_inject"`
	// 表示当前镜像所属的企业项目。 取值为0或无该值，表示属于default企业项目。 取值为UUID，表示属于该UUID对应的企业项目。 关于企业项目ID的获取及企业项目特性的详细信息，请参考《企业管理用户指南》。

	EnterpriseProjectId string `json:"enterprise_project_id"`
	// 云主机云服务器的启动方式。目前支持： bios：表示bios引导启动。 uefi：表示uefi引导启动。

	HwFirmwareType GlanceShowImageResponseBodyHwFirmwareType `json:"hw_firmware_type"`
	// 是否为ARM架构类型的镜像，取值为“true”或者“false”。

	SupportArm GlanceShowImageResponseBodySupportArm `json:"__support_arm"`
	// 表示当前市场镜像是否下架。 true：已下架 false：未下架

	IsOffshelved GlanceShowImageResponseBodyIsOffshelved `json:"__is_offshelved"`
}

func (o GlanceShowImageResponseBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceShowImageResponseBody struct{}"
	}

	return strings.Join([]string{"GlanceShowImageResponseBody", string(data)}, " ")
}

type GlanceShowImageResponseBodyImageSourceType struct {
	value string
}

type GlanceShowImageResponseBodyImageSourceTypeEnum struct {
	UDS   GlanceShowImageResponseBodyImageSourceType
	SWIFT GlanceShowImageResponseBodyImageSourceType
}

func GetGlanceShowImageResponseBodyImageSourceTypeEnum() GlanceShowImageResponseBodyImageSourceTypeEnum {
	return GlanceShowImageResponseBodyImageSourceTypeEnum{
		UDS: GlanceShowImageResponseBodyImageSourceType{
			value: "uds",
		},
		SWIFT: GlanceShowImageResponseBodyImageSourceType{
			value: "swift",
		},
	}
}

func (c GlanceShowImageResponseBodyImageSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyImageSourceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GlanceShowImageResponseBodyImagetype struct {
	value string
}

type GlanceShowImageResponseBodyImagetypeEnum struct {
	GOLD    GlanceShowImageResponseBodyImagetype
	PRIVATE GlanceShowImageResponseBodyImagetype
	SHARED  GlanceShowImageResponseBodyImagetype
}

func GetGlanceShowImageResponseBodyImagetypeEnum() GlanceShowImageResponseBodyImagetypeEnum {
	return GlanceShowImageResponseBodyImagetypeEnum{
		GOLD: GlanceShowImageResponseBodyImagetype{
			value: "gold",
		},
		PRIVATE: GlanceShowImageResponseBodyImagetype{
			value: "private",
		},
		SHARED: GlanceShowImageResponseBodyImagetype{
			value: "shared",
		},
	}
}

func (c GlanceShowImageResponseBodyImagetype) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyImagetype) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GlanceShowImageResponseBodyIsregistered struct {
	value string
}

type GlanceShowImageResponseBodyIsregisteredEnum struct {
	TRUE  GlanceShowImageResponseBodyIsregistered
	FALSE GlanceShowImageResponseBodyIsregistered
}

func GetGlanceShowImageResponseBodyIsregisteredEnum() GlanceShowImageResponseBodyIsregisteredEnum {
	return GlanceShowImageResponseBodyIsregisteredEnum{
		TRUE: GlanceShowImageResponseBodyIsregistered{
			value: "true",
		},
		FALSE: GlanceShowImageResponseBodyIsregistered{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseBodyIsregistered) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyIsregistered) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GlanceShowImageResponseBodyOsBit struct {
	value string
}

type GlanceShowImageResponseBodyOsBitEnum struct {
	E_32 GlanceShowImageResponseBodyOsBit
	E_64 GlanceShowImageResponseBodyOsBit
}

func GetGlanceShowImageResponseBodyOsBitEnum() GlanceShowImageResponseBodyOsBitEnum {
	return GlanceShowImageResponseBodyOsBitEnum{
		E_32: GlanceShowImageResponseBodyOsBit{
			value: "32",
		},
		E_64: GlanceShowImageResponseBodyOsBit{
			value: "64",
		},
	}
}

func (c GlanceShowImageResponseBodyOsBit) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyOsBit) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GlanceShowImageResponseBodyOsType struct {
	value string
}

type GlanceShowImageResponseBodyOsTypeEnum struct {
	LINUX   GlanceShowImageResponseBodyOsType
	WINDOWS GlanceShowImageResponseBodyOsType
	OTHER   GlanceShowImageResponseBodyOsType
}

func GetGlanceShowImageResponseBodyOsTypeEnum() GlanceShowImageResponseBodyOsTypeEnum {
	return GlanceShowImageResponseBodyOsTypeEnum{
		LINUX: GlanceShowImageResponseBodyOsType{
			value: "Linux",
		},
		WINDOWS: GlanceShowImageResponseBodyOsType{
			value: "Windows",
		},
		OTHER: GlanceShowImageResponseBodyOsType{
			value: "Other",
		},
	}
}

func (c GlanceShowImageResponseBodyOsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyOsType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GlanceShowImageResponseBodyPlatform struct {
	value string
}

type GlanceShowImageResponseBodyPlatformEnum struct {
	WINDOWS      GlanceShowImageResponseBodyPlatform
	UBUNTU       GlanceShowImageResponseBodyPlatform
	RED_HAT      GlanceShowImageResponseBodyPlatform
	SUSE         GlanceShowImageResponseBodyPlatform
	CENT_OS      GlanceShowImageResponseBodyPlatform
	DEBIAN       GlanceShowImageResponseBodyPlatform
	OPEN_SUSE    GlanceShowImageResponseBodyPlatform
	ORACLE_LINUX GlanceShowImageResponseBodyPlatform
	FEDORA       GlanceShowImageResponseBodyPlatform
	OTHER        GlanceShowImageResponseBodyPlatform
	CORE_OS      GlanceShowImageResponseBodyPlatform
	EULER_OS     GlanceShowImageResponseBodyPlatform
}

func GetGlanceShowImageResponseBodyPlatformEnum() GlanceShowImageResponseBodyPlatformEnum {
	return GlanceShowImageResponseBodyPlatformEnum{
		WINDOWS: GlanceShowImageResponseBodyPlatform{
			value: "Windows",
		},
		UBUNTU: GlanceShowImageResponseBodyPlatform{
			value: "Ubuntu",
		},
		RED_HAT: GlanceShowImageResponseBodyPlatform{
			value: "RedHat",
		},
		SUSE: GlanceShowImageResponseBodyPlatform{
			value: "SUSE",
		},
		CENT_OS: GlanceShowImageResponseBodyPlatform{
			value: "CentOS",
		},
		DEBIAN: GlanceShowImageResponseBodyPlatform{
			value: "Debian",
		},
		OPEN_SUSE: GlanceShowImageResponseBodyPlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: GlanceShowImageResponseBodyPlatform{
			value: "OracleLinux",
		},
		FEDORA: GlanceShowImageResponseBodyPlatform{
			value: "Fedora",
		},
		OTHER: GlanceShowImageResponseBodyPlatform{
			value: "Other",
		},
		CORE_OS: GlanceShowImageResponseBodyPlatform{
			value: "CoreOS",
		},
		EULER_OS: GlanceShowImageResponseBodyPlatform{
			value: "EulerOS",
		},
	}
}

func (c GlanceShowImageResponseBodyPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyPlatform) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GlanceShowImageResponseBodyDiskFormat struct {
	value string
}

type GlanceShowImageResponseBodyDiskFormatEnum struct {
	VHD   GlanceShowImageResponseBodyDiskFormat
	ZVHD  GlanceShowImageResponseBodyDiskFormat
	RAW   GlanceShowImageResponseBodyDiskFormat
	QCOW2 GlanceShowImageResponseBodyDiskFormat
	ZVHD2 GlanceShowImageResponseBodyDiskFormat
}

func GetGlanceShowImageResponseBodyDiskFormatEnum() GlanceShowImageResponseBodyDiskFormatEnum {
	return GlanceShowImageResponseBodyDiskFormatEnum{
		VHD: GlanceShowImageResponseBodyDiskFormat{
			value: "vhd",
		},
		ZVHD: GlanceShowImageResponseBodyDiskFormat{
			value: "zvhd",
		},
		RAW: GlanceShowImageResponseBodyDiskFormat{
			value: "raw",
		},
		QCOW2: GlanceShowImageResponseBodyDiskFormat{
			value: "qcow2",
		},
		ZVHD2: GlanceShowImageResponseBodyDiskFormat{
			value: "zvhd2",
		},
	}
}

func (c GlanceShowImageResponseBodyDiskFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyDiskFormat) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GlanceShowImageResponseBodyStatus struct {
	value string
}

type GlanceShowImageResponseBodyStatusEnum struct {
	QUEUED  GlanceShowImageResponseBodyStatus
	SAVING  GlanceShowImageResponseBodyStatus
	DELETED GlanceShowImageResponseBodyStatus
	KILLED  GlanceShowImageResponseBodyStatus
	ACTIVE  GlanceShowImageResponseBodyStatus
}

func GetGlanceShowImageResponseBodyStatusEnum() GlanceShowImageResponseBodyStatusEnum {
	return GlanceShowImageResponseBodyStatusEnum{
		QUEUED: GlanceShowImageResponseBodyStatus{
			value: "queued",
		},
		SAVING: GlanceShowImageResponseBodyStatus{
			value: "saving",
		},
		DELETED: GlanceShowImageResponseBodyStatus{
			value: "deleted",
		},
		KILLED: GlanceShowImageResponseBodyStatus{
			value: "killed",
		},
		ACTIVE: GlanceShowImageResponseBodyStatus{
			value: "active",
		},
	}
}

func (c GlanceShowImageResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GlanceShowImageResponseBodyVirtualEnvType struct {
	value string
}

type GlanceShowImageResponseBodyVirtualEnvTypeEnum struct {
	FUSION_COMPUTE GlanceShowImageResponseBodyVirtualEnvType
	IRONIC         GlanceShowImageResponseBodyVirtualEnvType
	DATA_IMAGE     GlanceShowImageResponseBodyVirtualEnvType
}

func GetGlanceShowImageResponseBodyVirtualEnvTypeEnum() GlanceShowImageResponseBodyVirtualEnvTypeEnum {
	return GlanceShowImageResponseBodyVirtualEnvTypeEnum{
		FUSION_COMPUTE: GlanceShowImageResponseBodyVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: GlanceShowImageResponseBodyVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: GlanceShowImageResponseBodyVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c GlanceShowImageResponseBodyVirtualEnvType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyVirtualEnvType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GlanceShowImageResponseBodyVisibility struct {
	value string
}

type GlanceShowImageResponseBodyVisibilityEnum struct {
	PRIVATE GlanceShowImageResponseBodyVisibility
	PUBLIC  GlanceShowImageResponseBodyVisibility
	SHARED  GlanceShowImageResponseBodyVisibility
}

func GetGlanceShowImageResponseBodyVisibilityEnum() GlanceShowImageResponseBodyVisibilityEnum {
	return GlanceShowImageResponseBodyVisibilityEnum{
		PRIVATE: GlanceShowImageResponseBodyVisibility{
			value: "private",
		},
		PUBLIC: GlanceShowImageResponseBodyVisibility{
			value: "public",
		},
		SHARED: GlanceShowImageResponseBodyVisibility{
			value: "shared",
		},
	}
}

func (c GlanceShowImageResponseBodyVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyVisibility) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GlanceShowImageResponseBodySupportFcInject struct {
	value string
}

type GlanceShowImageResponseBodySupportFcInjectEnum struct {
	TRUE  GlanceShowImageResponseBodySupportFcInject
	FALSE GlanceShowImageResponseBodySupportFcInject
}

func GetGlanceShowImageResponseBodySupportFcInjectEnum() GlanceShowImageResponseBodySupportFcInjectEnum {
	return GlanceShowImageResponseBodySupportFcInjectEnum{
		TRUE: GlanceShowImageResponseBodySupportFcInject{
			value: "true",
		},
		FALSE: GlanceShowImageResponseBodySupportFcInject{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseBodySupportFcInject) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodySupportFcInject) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GlanceShowImageResponseBodyHwFirmwareType struct {
	value string
}

type GlanceShowImageResponseBodyHwFirmwareTypeEnum struct {
	BIOS GlanceShowImageResponseBodyHwFirmwareType
	UEFI GlanceShowImageResponseBodyHwFirmwareType
}

func GetGlanceShowImageResponseBodyHwFirmwareTypeEnum() GlanceShowImageResponseBodyHwFirmwareTypeEnum {
	return GlanceShowImageResponseBodyHwFirmwareTypeEnum{
		BIOS: GlanceShowImageResponseBodyHwFirmwareType{
			value: "bios",
		},
		UEFI: GlanceShowImageResponseBodyHwFirmwareType{
			value: "uefi",
		},
	}
}

func (c GlanceShowImageResponseBodyHwFirmwareType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyHwFirmwareType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GlanceShowImageResponseBodySupportArm struct {
	value string
}

type GlanceShowImageResponseBodySupportArmEnum struct {
	TRUE  GlanceShowImageResponseBodySupportArm
	FALSE GlanceShowImageResponseBodySupportArm
}

func GetGlanceShowImageResponseBodySupportArmEnum() GlanceShowImageResponseBodySupportArmEnum {
	return GlanceShowImageResponseBodySupportArmEnum{
		TRUE: GlanceShowImageResponseBodySupportArm{
			value: "true",
		},
		FALSE: GlanceShowImageResponseBodySupportArm{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseBodySupportArm) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodySupportArm) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type GlanceShowImageResponseBodyIsOffshelved struct {
	value string
}

type GlanceShowImageResponseBodyIsOffshelvedEnum struct {
	TRUE  GlanceShowImageResponseBodyIsOffshelved
	FALSE GlanceShowImageResponseBodyIsOffshelved
}

func GetGlanceShowImageResponseBodyIsOffshelvedEnum() GlanceShowImageResponseBodyIsOffshelvedEnum {
	return GlanceShowImageResponseBodyIsOffshelvedEnum{
		TRUE: GlanceShowImageResponseBodyIsOffshelved{
			value: "true",
		},
		FALSE: GlanceShowImageResponseBodyIsOffshelved{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseBodyIsOffshelved) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyIsOffshelved) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
