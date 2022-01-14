package codegen

import "encoding/json"

// ConfigContext ConfigContext
type ConfigContext struct {
	DataSource     *DataSourceConfig `json:"dataSource"`
	GlobalConfig   *GlobalConfig     `json:"globalConfig"`
	PackageConfig  *PackageConfig    `json:"packageConfig"`
	TemplateConfig *TemplateConfig   `json:"templateConfig"`
	StrategyConfig *StrategyConfig   `json:"strategyConfig"`
}

func (cc *ConfigContext) String() string {
	data, _ := json.MarshalIndent(cc, "", "  ")
	return string(data)
}

func (dsc *DataSourceConfig) String() string {
	data, _ := json.MarshalIndent(dsc, "", "  ")
	return string(data)
}

// DataSourceConfig DataSourceConfig
type DataSourceConfig struct {
	Typ      string `json:"typ"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

// GlobalConfig GlobalConfig
type GlobalConfig struct {
	FileOverride   bool   `json:"fileOverride"`
	DisableOpenDir bool   `json:"disableOpenDir"`
	OutputDir      string `json:"outputDir"`
	Author         string `json:"author"`
	EnableKotlin   bool   `json:"enableKotlin"`
	EnableSwagger  bool   `json:"enableSwagger"`
	DateType       string `json:"dateType"`
	CommentDate    string `json:"commentDate"`
}

// PathInfo PathInfo
type PathInfo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// PackageConfig PackageConfig
type PackageConfig struct {
	Parent      string     `json:"parent"`
	ModuleName  string     `json:"moduleName"`
	Entity      string     `json:"entity"`
	Service     string     `json:"service"`
	ServiceImpl string     `json:"serviceImpl"`
	Mapper      string     `json:"mapper"`
	MapperXml   string     `json:"mapperXml"`
	Controller  string     `json:"controller"`
	Other       string     `json:"other"`
	PathInfo    []PathInfo `json:"pathInfo"`
}

// TemplateConfig TemplateConfig
type TemplateConfig struct {
	DisableAll  bool   `json:"disableAll"`
	Disable     bool   `json:"disable"`
	Entity      string `json:"entity"`
	EntityKt    string `json:"entityKt"`
	Service     string `json:"service"`
	ServiceImpl string `json:"serviceImpl"`
	Mapper      string `json:"mapper"`
	MapperXml   string `json:"mapperXml"`
	Controller  string `json:"controller"`
}

// InjectionConfig InjectionConfig
type InjectionConfig struct {
}

// StrategyConfig StrategyConfig
type StrategyConfig struct {
	EnableCapitalMode bool        `json:"enableCapitalMode"`
	EnableSkipView    bool        `json:"enableSkipView"`
	DisableSqlFilter  bool        `json:"disableSqlFilter"`
	EnableSchema      bool        `json:"enableSchema"`
	AddIncludes       []string    `json:"addIncludes"`
	Entity            *Entity     `json:"entity"`
	Controller        *Controller `json:"controller"`
	Mapper            *Mapper     `json:"mapper"`
	Service           *Service    `json:"service"`
}

// FieldTypeKeyVal 自动填充key-value
type FieldTypeKeyVal struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Entity Entity
type Entity struct {
	SuperClass                 string            `json:"superClass"`
	DisableSerialVersionUID    bool              `json:"disableSerialVersionUID"`
	EnableColumnConstant       bool              `json:"enableColumnConstant"`
	EnableChainModel           bool              `json:"enableChainModel"`
	EnableLombok               bool              `json:"enableLombok"`
	EnableRemoveIsPrefix       bool              `json:"enableRemoveIsPrefix"`
	EnableTableFieldAnnotation bool              `json:"enableTableFieldAnnotation"`
	EnableActiveRecord         bool              `json:"enableActiveRecord"`
	VersionColumnName          string            `json:"versionColumnName"`
	VersionPropertyName        string            `json:"versionPropertyName"`
	LogicDeleteColumnName      string            `json:"logicDeleteColumnName"`
	LogicDeletePropertyName    string            `json:"logicDeletePropertyName"`
	Naming                     string            `json:"naming"`
	AddSuperEntityColumns      []string          `json:"addSuperEntityColumns"`
	AddIgnoreColumns           []string          `json:"addIgnoreColumns"`
	AddTableFills              []FieldTypeKeyVal `json:"addTableFills"`
	IdType                     string            `json:"idType"`
	FormatFileName             string            `json:"formatFileName"`
}

// Controller Controller
type Controller struct {
	SuperClass        string `json:"superClass"`
	EnableHyphenStyle bool   `json:"enableHyphenStyle"`
	EnableRestStyle   bool   `json:"enableRestStyle"`
	FormatFileName    string `json:"formatFileName"`
}

// Mapper  Mapper
type Mapper struct {
	SuperClass             string `json:"superClass"`
	EnableMapperAnnotation bool   `json:"enableMapperAnnotation"`
	FormatMapperFileName   string `json:"formatMapperFileName"`
	FormatXmlFileName      string `json:"formatXmlFileName"`
}

// DatabaseOptions DatabaseOptions
type DatabaseOptions struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// Service Service
type Service struct {
	SuperServiceClass         string `json:"superServiceClass"`
	SuperServiceImplClass     string `json:"superServiceImplClass"`
	FormatServiceFileName     string `json:"formatServiceFileName"`
	FormatServiceImplFileName string `json:"formatServiceImplFileName"`
}
