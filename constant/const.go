package constant

const (
	Lower              = "abcdefghijklmnopqrstuvwxyz"
	Upper              = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Number             = "0123456789"
	Symbol             = "!@#$%^&*()[]"
	Pool               = Lower + Upper + Number + Symbol
	CommonChineseWords = "的一是在不了有和人这中大为上个国我以要他时来用们生到作地于出就分对成会可主发年动同工也能下过子" +
		"说产种面而方后多定行学法所民得经十三之进着等部度家电力里如水化高自二理起小物现实加量都两体制机当使点从业本去把性好应" +
		"开它合还因由其些然前外天政四日那社义事平形相全表间样与关各重新线内数正心反你明看原又么利比或但质气第向道命此变条只没" +
		"结解问意建月公无系军很情者最立代想已通并提直题党程展五果料象员革位入常文总次品式活设及管特件长求老头基资边流路级少图山" +
		"统接知较将组见计别她手角期根论运农指几九区强放决西被干做必战先回则任取据处队南给色光门即保治北造百规热领七海口东导器压" +
		"志世金增争济阶油思术极交受联什认六共权收证改清己美再采转更单风切打白教速花带安场身车例真务具万每目至达走积示议声报斗" +
		"完类八离华名确才科张信马节话米整空元况今集温传土许步群广石记需段研界拉林律叫且究观越织装影算低持音众书布复容儿须际商非验" +
		"连断深难近矿千周委素技备半办青省列习响约支般史感劳便团往酸历市克何除消构府称太准精值号率族维划选标写存候毛亲快效斯院查江" +
		"型眼王按格养易置派层片始却专状育厂京识适属圆包火住调满县局照参红细引听该铁价严龙飞"
)

var (
	First = []interface{}{
		"James", "John", "Robert", "Michael", "William",
		"David", "Richard", "Charles", "Joseph", "Thomas",
		"Christopher", "Daniel", "Paul", "Mark", "Donald",
		"George", "Kenneth", "Steven", "Edward", "Brian",
		"Ronald", "Anthony", "Kevin", "Jason", "Matthew",
		"Gary", "Timothy", "Jose", "Larry", "Jeffrey",
		"Frank", "Scott", "Eric", "Mary", "Patricia",
		"Linda", "Barbara", "Elizabeth", "Brenda", "Amy",
		"Jennifer", "Maria", "Susan", "Margaret", "Dorothy",
		"Lisa", "Nancy", "Karen", "Betty", "Helen",
		"Sandra", "Donna", "Carol", "Ruth", "Sharon",
		"Michelle", "Laura", "Sarah", "Kimberly", "Deborah",
		"Jessica", "Shirley", "Cynthia", "Angela", "Melissa",
		"Anna",
	}

	Last = []interface{}{
		"Smith", "Johnson", "Williams", "Brown", "Jones",
		"Miller", "Davis", "Garcia", "Rodriguez", "Wilson",
		"Martinez", "Anderson", "Taylor", "Thomas", "Hernandez",
		"Moore", "Martin", "Jackson", "Thompson", "White",
		"Lopez", "Lee", "Gonzalez", "Harris", "Clark",
		"Lewis", "Robinson", "Walker", "Perez", "Hall",
		"Young", "Allen",
	}

	CFirst = []interface{}{
		"王", "李", "张", "刘", "陈", "杨", "赵", "黄", "周", "吴",
		"徐", "孙", "胡", "朱", "高", "林", "何", "郭", "马", "罗",
		"梁", "宋", "郑", "谢", "韩", "唐", "冯", "于", "董", "萧",
		"程", "曹", "袁", "邓", "许", "傅", "沈", "曾", "彭", "吕",
		"苏", "卢", "蒋", "蔡", "贾", "丁", "魏", "薛", "叶", "阎",
		"余", "潘", "杜", "戴", "夏", "锺", "汪", "田", "任", "姜",
		"范", "方", "石", "姚", "谭", "廖", "邹", "熊", "金", "陆",
		"郝", "孔", "白", "崔", "康", "毛", "邱", "秦", "江", "史",
		"顾", "侯", "邵", "孟", "龙", "万", "段", "雷", "钱", "汤",
		"尹", "黎", "易", "常", "武", "乔", "贺", "赖", "龚", "文",
	}

	CLast = []interface{}{
		"伟", "芳", "娜", "秀", "英", "敏", "静", "丽", "强", "磊", "军",
		"洋", "勇", "艳", "杰", "娟", "涛", "明", "超", "秀", "兰", "霞",
		"平", "刚", "桂", "英",
	}

	Region = []interface{}{
		"东北", "华北", "华东", "华中",
		"华南", "西南", "西北",
	}

	Protocol = []interface{}{
		"http", "ftp", "gopher", "mailto", "mid", "cid", "news",
		"nntp", "prospero", "telnet", "rlogin", "tn3270", "wais",
	}

	TLD = []interface{}{
		// 域名后缀
		"com", "net", "org", "edu", "gov", "int", "mil", "cn",
		// 国内域名
		"com.cn", "net.cn", "gov.cn", "org.cn",
		// 中文国内域名
		"中国", "中国互联.公司", "中国互联.网络",
		// 新国际域名
		"tel", "biz", "cc", "tv", "info", "name", "hk", "mobi",
		"asia", "cd", "travel", "pro", "museum", "coop", "aero",
		// 世界各国域名后缀
		"ad", "ae", "af", "ag", "ai", "al", "am", "an", "ao",
		"aq", "ar", "as", "at", "au", "aw", "az", "ba", "bb",
		"bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn",
		"bo", "br", "bs", "bt", "bv", "bw", "by", "bz", "ca",
		"cc", "cf", "cg", "ch", "ci", "ck", "cl", "cm", "cn",
		"co", "cq", "cr", "cu", "cv", "cx", "cy", "cz", "de",
		"dj", "dk", "dm", "do", "dz", "ec", "ee", "eg", "eh",
		"es", "et", "ev", "fi", "fj", "fk", "fm", "fo", "fr",
		"ga", "gb", "gd", "ge", "gf", "gh", "gi", "gl", "gm",
		"gn", "gp", "gr", "gt", "gu", "gw", "gy", "hk", "hm",
		"hn", "hr", "ht", "hu", "id", "ie", "il", "in", "io",
		"iq", "ir", "is", "it", "jm", "jo", "jp", "ke", "kg",
		"kh", "ki", "km", "kn", "kp", "kr", "kw", "ky", "kz",
		"la", "lb", "lc", "li", "lk", "lr", "ls", "lt", "lu",
		"lv", "ly", "ma", "mc", "md", "mg", "mh", "ml", "mm",
		"mn", "mo", "mp", "mq", "mr", "ms", "mt", "mv", "mw",
		"mx", "my", "mz", "na", "nc", "ne", "nf", "ng", "ni",
		"nl", "no", "np", "nr", "nt", "nu", "nz", "om", "qa",
		"pa", "pe", "pf", "pg", "ph", "pk", "pl", "pm", "pn",
		"pr", "pt", "pw", "py", "re", "ro", "ru", "rw", "sa",
		"sb", "sc", "sd", "se", "sg", "sh", "si", "sj", "sk",
		"sl", "sm", "sn", "so", "sr", "st", "su", "sy", "sz",
		"tc", "td", "tf", "tg", "th", "tj", "tk", "tm", "tn",
		"to", "tp", "tr", "tt", "tv", "tw", "tz", "ua", "ug",
		"uk", "us", "uy", "va", "vc", "ve", "vg", "vn", "vu",
		"wf", "ws", "ye", "yu", "za", "zm", "zr", "zw",
	}
)
