package variables

const (
	VariableTypeString = "type.string"
	VariableTypeNumber = "type.number"
	VariableTypeBool   = "type.bool"
	VariableString     = "@string"
	VariableInteger    = "@integer"
	VariableFloat      = "@float"
	VariableCharacter  = "@character"
	VariableBool       = "@bool"
	VariableBoolean    = "@boolean"
	VariableNatural    = "@natural"
	VariableDate       = "@date"
	VariableTime       = "@time"
	VariableDateTime   = "@datetime"
	VariableNow        = "@now"
	VariableParagraph  = "@paragraph"
	VariableSentence   = "@sentence"
	VariableWord       = "@word"
	VariableTitle      = "@title"
	VariableCParagraph = "@cparagraph"
	VariableCSentence  = "@csentence"
	VariableCWord      = "@cword"
	VariableCTitle     = "@ctitle"
	VariableFirst      = "@first"
	VariableLast       = "@last"
	VariableName       = "@name"
	VariableCFirst     = "@cfirst"
	VariableCLast      = "@clast"
	VariableCname      = "@cname"
	VariableUrl        = "@url"
	VariableDomain     = "@domain"
	VariableProtocol   = "@protocol"
	VariableTld        = "@tld"
	VariableEmail      = "@email"
	VariableIp         = "@ip"
	VariableRegion     = "@region"
	VariableProvince   = "@province"
	VariableCity       = "@city"
	VariableCounty     = "@county"
	VariableZip        = "@zip"
	VariableCapitalize = "@capitalize"
	VariableUpper      = "@upper"
	VariableLower      = "@lower"
	VariablePick       = "@pick"
	VariableShuffle    = "@shuffle"
	VariableGuid       = "@guid"
	VariableId         = "@id"
	VariableColor      = "@color"
	VariableHex        = "@hex"
	VariableRGB        = "@rgb"
	VariableRGBA       = "@rgba"
	VariableHSL        = "@hsl"
	VariableImage      = "@image"
	VariableDataImage  = "@dataImage"
	VariableNull       = "@null"
)

var Factory = &factory{
	creators: make(map[string]Creator),
}

func RegisterVariable(name string, creator Creator) error {
	return Factory.Register(name, creator)
}

func init() {
	_ = RegisterVariable(VariableTypeString, DefaultRuleParserVariable(TypeStringFunc))
	_ = RegisterVariable(VariableTypeNumber, DefaultRuleParserVariable(TypeNumberFunc))
	_ = RegisterVariable(VariableString, DefaultRuleParserVariable(StringFunc))
	_ = RegisterVariable(VariableInteger, DefaultRuleParserVariable(IntegerFunc))
	_ = RegisterVariable(VariableFloat, DefaultRuleParserVariable(FloatFunc))
	_ = RegisterVariable(VariableCharacter, DefaultRuleParserVariable(CharacterFunc))
	_ = RegisterVariable(VariableBool, DefaultRuleParserVariable(BoolFunc))
	_ = RegisterVariable(VariableBoolean, DefaultRuleParserVariable(BoolFunc))
	_ = RegisterVariable(VariableTypeBool, DefaultRuleParserVariable(BoolFunc))
	_ = RegisterVariable(VariableNatural, DefaultRuleParserVariable(NaturalFunc))
	_ = RegisterVariable(VariableDate, DefaultRuleParserVariable(DateFunc))
	_ = RegisterVariable(VariableTime, DefaultRuleParserVariable(TimeFunc))
	_ = RegisterVariable(VariableDateTime, DefaultRuleParserVariable(DatetimeFunc))
	_ = RegisterVariable(VariableNow, DefaultRuleParserVariable(NowFunc))
	_ = RegisterVariable(VariableRegion, DefaultRuleParserVariable(RegionFunc))
	_ = RegisterVariable(VariableProvince, DefaultRuleParserVariable(ProvinceFunc))
	_ = RegisterVariable(VariableCity, DefaultRuleParserVariable(CityFunc))
	_ = RegisterVariable(VariableCounty, DefaultRuleParserVariable(CountryFunc))
	_ = RegisterVariable(VariableWord, DefaultRuleParserVariable(WordFunc))
	_ = RegisterVariable(VariableTitle, DefaultRuleParserVariable(TitleFunc))
	_ = RegisterVariable(VariableSentence, DefaultRuleParserVariable(SentenceFunc))
	_ = RegisterVariable(VariableParagraph, DefaultRuleParserVariable(ParagraphFunc))
	_ = RegisterVariable(VariableCWord, DefaultRuleParserVariable(CWordFunc))
	_ = RegisterVariable(VariableCTitle, DefaultRuleParserVariable(CTitleFunc))
	_ = RegisterVariable(VariableCSentence, DefaultRuleParserVariable(CSentenceFunc))
	_ = RegisterVariable(VariableCParagraph, DefaultRuleParserVariable(CParagraphFunc))
	_ = RegisterVariable(VariableFirst, DefaultRuleParserVariable(FirstFunc))
	_ = RegisterVariable(VariableLast, DefaultRuleParserVariable(LastFunc))
	_ = RegisterVariable(VariableName, DefaultRuleParserVariable(NameFunc))
	_ = RegisterVariable(VariableCFirst, DefaultRuleParserVariable(CFirstFunc))
	_ = RegisterVariable(VariableCLast, DefaultRuleParserVariable(CLastFunc))
	_ = RegisterVariable(VariableCname, DefaultRuleParserVariable(CNameFunc))
	_ = RegisterVariable(VariableUrl, DefaultRuleParserVariable(URLFunc))
	_ = RegisterVariable(VariableDomain, DefaultRuleParserVariable(DomainFunc))
	_ = RegisterVariable(VariableProtocol, DefaultRuleParserVariable(ProtocolFunc))
	_ = RegisterVariable(VariableTld, DefaultRuleParserVariable(TLDFunc))
	_ = RegisterVariable(VariableEmail, DefaultRuleParserVariable(EmailFunc))
	_ = RegisterVariable(VariableIp, DefaultRuleParserVariable(IPFunc))
	_ = RegisterVariable(VariableZip, DefaultRuleParserVariable(ZIPFunc))
	_ = RegisterVariable(VariableCapitalize, DefaultRuleParserVariable(CapitalizeFunc))
	_ = RegisterVariable(VariableUpper, DefaultRuleParserVariable(UpperFunc))
	_ = RegisterVariable(VariableLower, DefaultRuleParserVariable(LowerFunc))
	_ = RegisterVariable(VariablePick, DefaultRuleParserVariable(PickFunc))
	_ = RegisterVariable(VariableShuffle, DefaultRuleParserVariable(ShuffleFunc))
	_ = RegisterVariable(VariableGuid, DefaultRuleParserVariable(GUIDFunc))
	_ = RegisterVariable(VariableId, DefaultRuleParserVariable(IDFunc))
	_ = RegisterVariable(VariableColor, DefaultRuleParserVariable(ColorFunc))
	_ = RegisterVariable(VariableHex, DefaultRuleParserVariable(ColorFunc))
	_ = RegisterVariable(VariableRGB, DefaultRuleParserVariable(RGBFunc))
	_ = RegisterVariable(VariableRGBA, DefaultRuleParserVariable(RGBAFunc))
	_ = RegisterVariable(VariableHSL, DefaultRuleParserVariable(HSLFunc))
	_ = RegisterVariable(VariableImage, DefaultRuleParserVariable(ImageFunc))
	_ = RegisterVariable(VariableDataImage, DefaultRuleParserVariable(DataImageFunc))
	_ = RegisterVariable(VariableNull, DefaultRuleParserVariable(NullFunc))
}

// ======================
type ImageURL interface {
	GenImageURL(string) string
}

var imageAdaptor ImageURL

type ImageURLFunc func(string) string

func (f ImageURLFunc) GenImageURL(id string) string {
	return f(id)
}

func RegisterImageURLAdaptor(ia ImageURL) {
	imageAdaptor = ia
}
