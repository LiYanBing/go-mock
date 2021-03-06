package constant

var Provinces = []*Province{
	{
		Name: "北京",
		Children: []*City{
			{
				Name: "北京市",
				Children: []*Distinct{
					{
						Name: "东城区",
					},
					{
						Name: "西城区",
					},
					{
						Name: "朝阳区",
					},
					{
						Name: "丰台区",
					},
					{
						Name: "石景山区",
					},
					{
						Name: "海淀区",
					},
					{
						Name: "门头沟区",
					},
					{
						Name: "房山区",
					},
					{
						Name: "通州区",
					},
					{
						Name: "顺义区",
					},
					{
						Name: "昌平区",
					},
					{
						Name: "大兴区",
					},
					{
						Name: "怀柔区",
					},
					{
						Name: "平谷区",
					},
				},
			},
		},
	},
	{
		Name: "天津",
		Children: []*City{
			{
				Name: "天津市",
				Children: []*Distinct{
					{
						Name: "和平区",
					},
					{
						Name: "河东区",
					},
					{
						Name: "河西区",
					},
					{
						Name: "南开区",
					},
					{
						Name: "河北区",
					},
					{
						Name: "红桥区",
					},
					{
						Name: "东丽区",
					},
					{
						Name: "西青区",
					},
					{
						Name: "津南区",
					},
					{
						Name: "北辰区",
					},
					{
						Name: "武清区",
					},
					{
						Name: "宝坻区",
					},
					{
						Name: "滨海新区",
					},
				},
			},
		},
	},
	{
		Name: "河北省",
		Children: []*City{
			{
				Name: "石家庄市",
				Children: []*Distinct{
					{
						Name: "长安区",
					},
					{
						Name: "桥东区",
					},
					{
						Name: "桥西区",
					},
					{
						Name: "新华区",
					},
					{
						Name: "井陉矿区",
					},
					{
						Name: "裕华区",
					},
					{
						Name: "井陉县",
					},
					{
						Name: "正定县",
					},
					{
						Name: "栾城县",
					},
					{
						Name: "行唐县",
					},
					{
						Name: "灵寿县",
					},
					{
						Name: "高邑县",
					},
					{
						Name: "深泽县",
					},
					{
						Name: "赞皇县",
					},
					{
						Name: "无极县",
					},
					{
						Name: "平山县",
					},
					{
						Name: "元氏县",
					},
					{
						Name: "赵县",
					},
					{
						Name: "辛集市",
					},
					{
						Name: "藁城市",
					},
					{
						Name: "晋州市",
					},
					{
						Name: "新乐市",
					},
					{
						Name: "鹿泉市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "唐山市",
				Children: []*Distinct{
					{
						Name: "路南区",
					},
					{
						Name: "路北区",
					},
					{
						Name: "古冶区",
					},
					{
						Name: "开平区",
					},
					{
						Name: "丰南区",
					},
					{
						Name: "丰润区",
					},
					{
						Name: "滦县",
					},
					{
						Name: "滦南县",
					},
					{
						Name: "乐亭县",
					},
					{
						Name: "迁西县",
					},
					{
						Name: "玉田县",
					},
					{
						Name: "曹妃甸区",
					},
					{
						Name: "遵化市",
					},
					{
						Name: "迁安市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "秦皇岛市",
				Children: []*Distinct{
					{
						Name: "海港区",
					},
					{
						Name: "山海关区",
					},
					{
						Name: "北戴河区",
					},
					{
						Name: "青龙满族自治县",
					},
					{
						Name: "昌黎县",
					},
					{
						Name: "抚宁县",
					},
					{
						Name: "卢龙县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "邯郸市",
				Children: []*Distinct{
					{
						Name: "邯山区",
					},
					{
						Name: "丛台区",
					},
					{
						Name: "复兴区",
					},
					{
						Name: "峰峰矿区",
					},
					{
						Name: "邯郸县",
					},
					{
						Name: "临漳县",
					},
					{
						Name: "成安县",
					},
					{
						Name: "大名县",
					},
					{
						Name: "涉县",
					},
					{
						Name: "磁县",
					},
					{
						Name: "肥乡县",
					},
					{
						Name: "永年县",
					},
					{
						Name: "邱县",
					},
					{
						Name: "鸡泽县",
					},
					{
						Name: "广平县",
					},
					{
						Name: "馆陶县",
					},
					{
						Name: "魏县",
					},
					{
						Name: "曲周县",
					},
					{
						Name: "武安市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "邢台市",
				Children: []*Distinct{
					{
						Name: "桥东区",
					},
					{
						Name: "桥西区",
					},
					{
						Name: "邢台县",
					},
					{
						Name: "临城县",
					},
					{
						Name: "内丘县",
					},
					{
						Name: "柏乡县",
					},
					{
						Name: "隆尧县",
					},
					{
						Name: "任县",
					},
					{
						Name: "南和县",
					},
					{
						Name: "宁晋县",
					},
					{
						Name: "巨鹿县",
					},
					{
						Name: "新河县",
					},
					{
						Name: "广宗县",
					},
					{
						Name: "平乡县",
					},
					{
						Name: "威县",
					},
					{
						Name: "清河县",
					},
					{
						Name: "临西县",
					},
					{
						Name: "南宫市",
					},
					{
						Name: "沙河市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "保定市",
				Children: []*Distinct{
					{
						Name: "新市区",
					},
					{
						Name: "北市区",
					},
					{
						Name: "南市区",
					},
					{
						Name: "满城县",
					},
					{
						Name: "清苑县",
					},
					{
						Name: "涞水县",
					},
					{
						Name: "阜平县",
					},
					{
						Name: "徐水县",
					},
					{
						Name: "定兴县",
					},
					{
						Name: "唐县",
					},
					{
						Name: "高阳县",
					},
					{
						Name: "容城县",
					},
					{
						Name: "涞源县",
					},
					{
						Name: "望都县",
					},
					{
						Name: "安新县",
					},
					{
						Name: "易县",
					},
					{
						Name: "曲阳县",
					},
					{
						Name: "蠡县",
					},
					{
						Name: "顺平县",
					},
					{
						Name: "博野县",
					},
					{
						Name: "雄县",
					},
					{
						Name: "涿州市",
					},
					{
						Name: "定州市",
					},
					{
						Name: "安国市",
					},
					{
						Name: "高碑店市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "张家口市",
				Children: []*Distinct{
					{
						Name: "桥东区",
					},
					{
						Name: "桥西区",
					},
					{
						Name: "宣化区",
					},
					{
						Name: "下花园区",
					},
					{
						Name: "宣化县",
					},
					{
						Name: "张北县",
					},
					{
						Name: "康保县",
					},
					{
						Name: "沽源县",
					},
					{
						Name: "尚义县",
					},
					{
						Name: "蔚县",
					},
					{
						Name: "阳原县",
					},
					{
						Name: "怀安县",
					},
					{
						Name: "万全县",
					},
					{
						Name: "怀来县",
					},
					{
						Name: "涿鹿县",
					},
					{
						Name: "赤城县",
					},
					{
						Name: "崇礼县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "承德市",
				Children: []*Distinct{
					{
						Name: "双桥区",
					},
					{
						Name: "双滦区",
					},
					{
						Name: "鹰手营子矿区",
					},
					{
						Name: "承德县",
					},
					{
						Name: "兴隆县",
					},
					{
						Name: "平泉县",
					},
					{
						Name: "滦平县",
					},
					{
						Name: "隆化县",
					},
					{
						Name: "丰宁满族自治县",
					},
					{
						Name: "宽城满族自治县",
					},
					{
						Name: "围场满族蒙古族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "沧州市",
				Children: []*Distinct{
					{
						Name: "新华区",
					},
					{
						Name: "运河区",
					},
					{
						Name: "沧县",
					},
					{
						Name: "青县",
					},
					{
						Name: "东光县",
					},
					{
						Name: "海兴县",
					},
					{
						Name: "盐山县",
					},
					{
						Name: "肃宁县",
					},
					{
						Name: "南皮县",
					},
					{
						Name: "吴桥县",
					},
					{
						Name: "献县",
					},
					{
						Name: "孟村回族自治县",
					},
					{
						Name: "泊头市",
					},
					{
						Name: "任丘市",
					},
					{
						Name: "黄骅市",
					},
					{
						Name: "河间市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "廊坊市",
				Children: []*Distinct{
					{
						Name: "安次区",
					},
					{
						Name: "广阳区",
					},
					{
						Name: "固安县",
					},
					{
						Name: "永清县",
					},
					{
						Name: "香河县",
					},
					{
						Name: "大城县",
					},
					{
						Name: "文安县",
					},
					{
						Name: "大厂回族自治县",
					},
					{
						Name: "霸州市",
					},
					{
						Name: "三河市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "衡水市",
				Children: []*Distinct{
					{
						Name: "桃城区",
					},
					{
						Name: "枣强县",
					},
					{
						Name: "武邑县",
					},
					{
						Name: "武强县",
					},
					{
						Name: "饶阳县",
					},
					{
						Name: "安平县",
					},
					{
						Name: "故城县",
					},
					{
						Name: "景县",
					},
					{
						Name: "阜城县",
					},
					{
						Name: "冀州市",
					},
					{
						Name: "深州市",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "山西省",
		Children: []*City{
			{
				Name: "太原市",
				Children: []*Distinct{
					{
						Name: "小店区",
					},
					{
						Name: "迎泽区",
					},
					{
						Name: "杏花岭区",
					},
					{
						Name: "尖草坪区",
					},
					{
						Name: "万柏林区",
					},
					{
						Name: "晋源区",
					},
					{
						Name: "清徐县",
					},
					{
						Name: "阳曲县",
					},
					{
						Name: "娄烦县",
					},
					{
						Name: "古交市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "大同市",
				Children: []*Distinct{
					{
						Name: "城区",
					},
					{
						Name: "矿区",
					},
					{
						Name: "南郊区",
					},
					{
						Name: "新荣区",
					},
					{
						Name: "阳高县",
					},
					{
						Name: "天镇县",
					},
					{
						Name: "广灵县",
					},
					{
						Name: "灵丘县",
					},
					{
						Name: "浑源县",
					},
					{
						Name: "左云县",
					},
					{
						Name: "大同县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "阳泉市",
				Children: []*Distinct{
					{
						Name: "城区",
					},
					{
						Name: "矿区",
					},
					{
						Name: "郊区",
					},
					{
						Name: "平定县",
					},
					{
						Name: "盂县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "长治市",
				Children: []*Distinct{
					{
						Name: "长治县",
					},
					{
						Name: "襄垣县",
					},
					{
						Name: "屯留县",
					},
					{
						Name: "平顺县",
					},
					{
						Name: "黎城县",
					},
					{
						Name: "壶关县",
					},
					{
						Name: "长子县",
					},
					{
						Name: "武乡县",
					},
					{
						Name: "沁县",
					},
					{
						Name: "沁源县",
					},
					{
						Name: "潞城市",
					},
					{
						Name: "城区",
					},
					{
						Name: "郊区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "晋城市",
				Children: []*Distinct{
					{
						Name: "城区",
					},
					{
						Name: "沁水县",
					},
					{
						Name: "阳城县",
					},
					{
						Name: "陵川县",
					},
					{
						Name: "泽州县",
					},
					{
						Name: "高平市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "朔州市",
				Children: []*Distinct{
					{
						Name: "朔城区",
					},
					{
						Name: "平鲁区",
					},
					{
						Name: "山阴县",
					},
					{
						Name: "应县",
					},
					{
						Name: "右玉县",
					},
					{
						Name: "怀仁县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "晋中市",
				Children: []*Distinct{
					{
						Name: "榆次区",
					},
					{
						Name: "榆社县",
					},
					{
						Name: "左权县",
					},
					{
						Name: "和顺县",
					},
					{
						Name: "昔阳县",
					},
					{
						Name: "寿阳县",
					},
					{
						Name: "太谷县",
					},
					{
						Name: "祁县",
					},
					{
						Name: "平遥县",
					},
					{
						Name: "灵石县",
					},
					{
						Name: "介休市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "运城市",
				Children: []*Distinct{
					{
						Name: "盐湖区",
					},
					{
						Name: "临猗县",
					},
					{
						Name: "万荣县",
					},
					{
						Name: "闻喜县",
					},
					{
						Name: "稷山县",
					},
					{
						Name: "新绛县",
					},
					{
						Name: "绛县",
					},
					{
						Name: "垣曲县",
					},
					{
						Name: "夏县",
					},
					{
						Name: "平陆县",
					},
					{
						Name: "芮城县",
					},
					{
						Name: "永济市",
					},
					{
						Name: "河津市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "忻州市",
				Children: []*Distinct{
					{
						Name: "忻府区",
					},
					{
						Name: "定襄县",
					},
					{
						Name: "五台县",
					},
					{
						Name: "代县",
					},
					{
						Name: "繁峙县",
					},
					{
						Name: "宁武县",
					},
					{
						Name: "静乐县",
					},
					{
						Name: "神池县",
					},
					{
						Name: "五寨县",
					},
					{
						Name: "岢岚县",
					},
					{
						Name: "河曲县",
					},
					{
						Name: "保德县",
					},
					{
						Name: "偏关县",
					},
					{
						Name: "原平市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "临汾市",
				Children: []*Distinct{
					{
						Name: "尧都区",
					},
					{
						Name: "曲沃县",
					},
					{
						Name: "翼城县",
					},
					{
						Name: "襄汾县",
					},
					{
						Name: "洪洞县",
					},
					{
						Name: "古县",
					},
					{
						Name: "安泽县",
					},
					{
						Name: "浮山县",
					},
					{
						Name: "吉县",
					},
					{
						Name: "乡宁县",
					},
					{
						Name: "大宁县",
					},
					{
						Name: "隰县",
					},
					{
						Name: "永和县",
					},
					{
						Name: "蒲县",
					},
					{
						Name: "汾西县",
					},
					{
						Name: "侯马市",
					},
					{
						Name: "霍州市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "吕梁市",
				Children: []*Distinct{
					{
						Name: "离石区",
					},
					{
						Name: "文水县",
					},
					{
						Name: "交城县",
					},
					{
						Name: "兴县",
					},
					{
						Name: "临县",
					},
					{
						Name: "柳林县",
					},
					{
						Name: "石楼县",
					},
					{
						Name: "岚县",
					},
					{
						Name: "方山县",
					},
					{
						Name: "中阳县",
					},
					{
						Name: "交口县",
					},
					{
						Name: "孝义市",
					},
					{
						Name: "汾阳市",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "内蒙古自治区",
		Children: []*City{
			{
				Name: "呼和浩特市",
				Children: []*Distinct{
					{
						Name: "新城区",
					},
					{
						Name: "回民区",
					},
					{
						Name: "玉泉区",
					},
					{
						Name: "赛罕区",
					},
					{
						Name: "土默特左旗",
					},
					{
						Name: "托克托县",
					},
					{
						Name: "和林格尔县",
					},
					{
						Name: "清水河县",
					},
					{
						Name: "武川县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "包头市",
				Children: []*Distinct{
					{
						Name: "东河区",
					},
					{
						Name: "昆都仑区",
					},
					{
						Name: "青山区",
					},
					{
						Name: "石拐区",
					},
					{
						Name: "白云鄂博矿区",
					},
					{
						Name: "九原区",
					},
					{
						Name: "土默特右旗",
					},
					{
						Name: "固阳县",
					},
					{
						Name: "达尔罕茂明安联合旗",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "乌海市",
				Children: []*Distinct{
					{
						Name: "海勃湾区",
					},
					{
						Name: "海南区",
					},
					{
						Name: "乌达区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "赤峰市",
				Children: []*Distinct{
					{
						Name: "红山区",
					},
					{
						Name: "元宝山区",
					},
					{
						Name: "松山区",
					},
					{
						Name: "阿鲁科尔沁旗",
					},
					{
						Name: "巴林左旗",
					},
					{
						Name: "巴林右旗",
					},
					{
						Name: "林西县",
					},
					{
						Name: "克什克腾旗",
					},
					{
						Name: "翁牛特旗",
					},
					{
						Name: "喀喇沁旗",
					},
					{
						Name: "宁城县",
					},
					{
						Name: "敖汉旗",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "通辽市",
				Children: []*Distinct{
					{
						Name: "科尔沁区",
					},
					{
						Name: "科尔沁左翼中旗",
					},
					{
						Name: "科尔沁左翼后旗",
					},
					{
						Name: "开鲁县",
					},
					{
						Name: "库伦旗",
					},
					{
						Name: "奈曼旗",
					},
					{
						Name: "扎鲁特旗",
					},
					{
						Name: "霍林郭勒市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "鄂尔多斯市",
				Children: []*Distinct{
					{
						Name: "东胜区",
					},
					{
						Name: "达拉特旗",
					},
					{
						Name: "准格尔旗",
					},
					{
						Name: "鄂托克前旗",
					},
					{
						Name: "鄂托克旗",
					},
					{
						Name: "杭锦旗",
					},
					{
						Name: "乌审旗",
					},
					{
						Name: "伊金霍洛旗",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "呼伦贝尔市",
				Children: []*Distinct{
					{
						Name: "海拉尔区",
					},
					{
						Name: "扎赉诺尔区",
					},
					{
						Name: "阿荣旗",
					},
					{
						Name: "莫力达瓦达斡尔族自治旗",
					},
					{
						Name: "鄂伦春自治旗",
					},
					{
						Name: "鄂温克族自治旗",
					},
					{
						Name: "陈巴尔虎旗",
					},
					{
						Name: "新巴尔虎左旗",
					},
					{
						Name: "新巴尔虎右旗",
					},
					{
						Name: "满洲里市",
					},
					{
						Name: "牙克石市",
					},
					{
						Name: "扎兰屯市",
					},
					{
						Name: "额尔古纳市",
					},
					{
						Name: "根河市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "巴彦淖尔市",
				Children: []*Distinct{
					{
						Name: "临河区",
					},
					{
						Name: "五原县",
					},
					{
						Name: "磴口县",
					},
					{
						Name: "乌拉特前旗",
					},
					{
						Name: "乌拉特中旗",
					},
					{
						Name: "乌拉特后旗",
					},
					{
						Name: "杭锦后旗",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "乌兰察布市",
				Children: []*Distinct{
					{
						Name: "集宁区",
					},
					{
						Name: "卓资县",
					},
					{
						Name: "化德县",
					},
					{
						Name: "商都县",
					},
					{
						Name: "兴和县",
					},
					{
						Name: "凉城县",
					},
					{
						Name: "察哈尔右翼前旗",
					},
					{
						Name: "察哈尔右翼中旗",
					},
					{
						Name: "察哈尔右翼后旗",
					},
					{
						Name: "四子王旗",
					},
					{
						Name: "丰镇市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "兴安盟",
				Children: []*Distinct{
					{
						Name: "乌兰浩特市",
					},
					{
						Name: "阿尔山市",
					},
					{
						Name: "科尔沁右翼前旗",
					},
					{
						Name: "科尔沁右翼中旗",
					},
					{
						Name: "扎赉特旗",
					},
					{
						Name: "突泉县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "锡林郭勒盟",
				Children: []*Distinct{
					{
						Name: "二连浩特市",
					},
					{
						Name: "锡林浩特市",
					},
					{
						Name: "阿巴嘎旗",
					},
					{
						Name: "苏尼特左旗",
					},
					{
						Name: "苏尼特右旗",
					},
					{
						Name: "东乌珠穆沁旗",
					},
					{
						Name: "西乌珠穆沁旗",
					},
					{
						Name: "太仆寺旗",
					},
					{
						Name: "镶黄旗",
					},
					{
						Name: "正镶白旗",
					},
					{
						Name: "正蓝旗",
					},
					{
						Name: "多伦县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "阿拉善盟",
				Children: []*Distinct{
					{
						Name: "阿拉善左旗",
					},
					{
						Name: "阿拉善右旗",
					},
					{
						Name: "额济纳旗",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "辽宁省",
		Children: []*City{
			{
				Name: "沈阳市",
				Children: []*Distinct{
					{
						Name: "和平区",
					},
					{
						Name: "沈河区",
					},
					{
						Name: "大东区",
					},
					{
						Name: "皇姑区",
					},
					{
						Name: "铁西区",
					},
					{
						Name: "苏家屯区",
					},
					{
						Name: "东陵区",
					},
					{
						Name: "新城子区",
					},
					{
						Name: "于洪区",
					},
					{
						Name: "辽中县",
					},
					{
						Name: "康平县",
					},
					{
						Name: "法库县",
					},
					{
						Name: "新民市",
					},
					{
						Name: "沈北新区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "大连市",
				Children: []*Distinct{
					{
						Name: "中山区",
					},
					{
						Name: "西岗区",
					},
					{
						Name: "沙河口区",
					},
					{
						Name: "甘井子区",
					},
					{
						Name: "旅顺口区",
					},
					{
						Name: "金州区",
					},
					{
						Name: "长海县",
					},
					{
						Name: "瓦房店市",
					},
					{
						Name: "普兰店市",
					},
					{
						Name: "庄河市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "鞍山市",
				Children: []*Distinct{
					{
						Name: "铁东区",
					},
					{
						Name: "铁西区",
					},
					{
						Name: "立山区",
					},
					{
						Name: "千山区",
					},
					{
						Name: "台安县",
					},
					{
						Name: "岫岩满族自治县",
					},
					{
						Name: "海城市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "抚顺市",
				Children: []*Distinct{
					{
						Name: "新抚区",
					},
					{
						Name: "东洲区",
					},
					{
						Name: "望花区",
					},
					{
						Name: "顺城区",
					},
					{
						Name: "抚顺县",
					},
					{
						Name: "新宾满族自治县",
					},
					{
						Name: "清原满族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "本溪市",
				Children: []*Distinct{
					{
						Name: "平山区",
					},
					{
						Name: "溪湖区",
					},
					{
						Name: "明山区",
					},
					{
						Name: "南芬区",
					},
					{
						Name: "本溪满族自治县",
					},
					{
						Name: "桓仁满族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "丹东市",
				Children: []*Distinct{
					{
						Name: "元宝区",
					},
					{
						Name: "振兴区",
					},
					{
						Name: "振安区",
					},
					{
						Name: "宽甸满族自治县",
					},
					{
						Name: "东港市",
					},
					{
						Name: "凤城市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "锦州市",
				Children: []*Distinct{
					{
						Name: "古塔区",
					},
					{
						Name: "凌河区",
					},
					{
						Name: "太和区",
					},
					{
						Name: "黑山县",
					},
					{
						Name: "义县",
					},
					{
						Name: "凌海市",
					},
					{
						Name: "北镇市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "营口市",
				Children: []*Distinct{
					{
						Name: "站前区",
					},
					{
						Name: "西市区",
					},
					{
						Name: "鲅鱼圈区",
					},
					{
						Name: "老边区",
					},
					{
						Name: "盖州市",
					},
					{
						Name: "大石桥市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "阜新市",
				Children: []*Distinct{
					{
						Name: "海州区",
					},
					{
						Name: "新邱区",
					},
					{
						Name: "太平区",
					},
					{
						Name: "清河门区",
					},
					{
						Name: "细河区",
					},
					{
						Name: "阜新蒙古族自治县",
					},
					{
						Name: "彰武县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "辽阳市",
				Children: []*Distinct{
					{
						Name: "白塔区",
					},
					{
						Name: "文圣区",
					},
					{
						Name: "宏伟区",
					},
					{
						Name: "弓长岭区",
					},
					{
						Name: "太子河区",
					},
					{
						Name: "辽阳县",
					},
					{
						Name: "灯塔市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "盘锦市",
				Children: []*Distinct{
					{
						Name: "双台子区",
					},
					{
						Name: "兴隆台区",
					},
					{
						Name: "大洼县",
					},
					{
						Name: "盘山县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "铁岭市",
				Children: []*Distinct{
					{
						Name: "银州区",
					},
					{
						Name: "清河区",
					},
					{
						Name: "铁岭县",
					},
					{
						Name: "西丰县",
					},
					{
						Name: "昌图县",
					},
					{
						Name: "调兵山市",
					},
					{
						Name: "开原市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "朝阳市",
				Children: []*Distinct{
					{
						Name: "双塔区",
					},
					{
						Name: "龙城区",
					},
					{
						Name: "朝阳县",
					},
					{
						Name: "建平县",
					},
					{
						Name: "喀喇沁左翼蒙古族自治县",
					},
					{
						Name: "北票市",
					},
					{
						Name: "凌源市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "葫芦岛市",
				Children: []*Distinct{
					{
						Name: "连山区",
					},
					{
						Name: "龙港区",
					},
					{
						Name: "南票区",
					},
					{
						Name: "绥中县",
					},
					{
						Name: "建昌县",
					},
					{
						Name: "兴城市",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "吉林省",
		Children: []*City{
			{
				Name: "长春市",
				Children: []*Distinct{
					{
						Name: "南关区",
					},
					{
						Name: "宽城区",
					},
					{
						Name: "朝阳区",
					},
					{
						Name: "二道区",
					},
					{
						Name: "绿园区",
					},
					{
						Name: "双阳区",
					},
					{
						Name: "农安县",
					},
					{
						Name: "九台市",
					},
					{
						Name: "榆树市",
					},
					{
						Name: "德惠市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "吉林市",
				Children: []*Distinct{
					{
						Name: "昌邑区",
					},
					{
						Name: "龙潭区",
					},
					{
						Name: "船营区",
					},
					{
						Name: "丰满区",
					},
					{
						Name: "永吉县",
					},
					{
						Name: "蛟河市",
					},
					{
						Name: "桦甸市",
					},
					{
						Name: "舒兰市",
					},
					{
						Name: "磐石市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "四平市",
				Children: []*Distinct{
					{
						Name: "铁西区",
					},
					{
						Name: "铁东区",
					},
					{
						Name: "梨树县",
					},
					{
						Name: "伊通满族自治县",
					},
					{
						Name: "公主岭市",
					},
					{
						Name: "双辽市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "辽源市",
				Children: []*Distinct{
					{
						Name: "龙山区",
					},
					{
						Name: "西安区",
					},
					{
						Name: "东丰县",
					},
					{
						Name: "东辽县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "通化市",
				Children: []*Distinct{
					{
						Name: "东昌区",
					},
					{
						Name: "二道江区",
					},
					{
						Name: "通化县",
					},
					{
						Name: "辉南县",
					},
					{
						Name: "柳河县",
					},
					{
						Name: "梅河口市",
					},
					{
						Name: "集安市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "白山市",
				Children: []*Distinct{
					{
						Name: "浑江区",
					},
					{
						Name: "抚松县",
					},
					{
						Name: "靖宇县",
					},
					{
						Name: "长白朝鲜族自治县",
					},
					{
						Name: "江源区",
					},
					{
						Name: "临江市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "松原市",
				Children: []*Distinct{
					{
						Name: "宁江区",
					},
					{
						Name: "前郭尔罗斯蒙古族自治县",
					},
					{
						Name: "长岭县",
					},
					{
						Name: "乾安县",
					},
					{
						Name: "扶余市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "白城市",
				Children: []*Distinct{
					{
						Name: "洮北区",
					},
					{
						Name: "镇赉县",
					},
					{
						Name: "通榆县",
					},
					{
						Name: "洮南市",
					},
					{
						Name: "大安市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "延边朝鲜族自治州",
				Children: []*Distinct{
					{
						Name: "延吉市",
					},
					{
						Name: "图们市",
					},
					{
						Name: "敦化市",
					},
					{
						Name: "珲春市",
					},
					{
						Name: "龙井市",
					},
					{
						Name: "和龙市",
					},
					{
						Name: "汪清县",
					},
					{
						Name: "安图县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "黑龙江省",
		Children: []*City{
			{
				Name: "哈尔滨市",
				Children: []*Distinct{
					{
						Name: "道里区",
					},
					{
						Name: "南岗区",
					},
					{
						Name: "道外区",
					},
					{
						Name: "香坊区",
					},
					{
						Name: "平房区",
					},
					{
						Name: "松北区",
					},
					{
						Name: "呼兰区",
					},
					{
						Name: "依兰县",
					},
					{
						Name: "方正县",
					},
					{
						Name: "宾县",
					},
					{
						Name: "巴彦县",
					},
					{
						Name: "木兰县",
					},
					{
						Name: "通河县",
					},
					{
						Name: "延寿县",
					},
					{
						Name: "阿城区",
					},
					{
						Name: "双城市",
					},
					{
						Name: "尚志市",
					},
					{
						Name: "五常市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "齐齐哈尔市",
				Children: []*Distinct{
					{
						Name: "龙沙区",
					},
					{
						Name: "建华区",
					},
					{
						Name: "铁锋区",
					},
					{
						Name: "昂昂溪区",
					},
					{
						Name: "富拉尔基区",
					},
					{
						Name: "碾子山区",
					},
					{
						Name: "梅里斯达斡尔族区",
					},
					{
						Name: "龙江县",
					},
					{
						Name: "依安县",
					},
					{
						Name: "泰来县",
					},
					{
						Name: "甘南县",
					},
					{
						Name: "富裕县",
					},
					{
						Name: "克山县",
					},
					{
						Name: "克东县",
					},
					{
						Name: "拜泉县",
					},
					{
						Name: "讷河市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "鸡西市",
				Children: []*Distinct{
					{
						Name: "鸡冠区",
					},
					{
						Name: "恒山区",
					},
					{
						Name: "滴道区",
					},
					{
						Name: "梨树区",
					},
					{
						Name: "城子河区",
					},
					{
						Name: "麻山区",
					},
					{
						Name: "鸡东县",
					},
					{
						Name: "虎林市",
					},
					{
						Name: "密山市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "鹤岗市",
				Children: []*Distinct{
					{
						Name: "向阳区",
					},
					{
						Name: "工农区",
					},
					{
						Name: "南山区",
					},
					{
						Name: "兴安区",
					},
					{
						Name: "东山区",
					},
					{
						Name: "兴山区",
					},
					{
						Name: "萝北县",
					},
					{
						Name: "绥滨县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "双鸭山市",
				Children: []*Distinct{
					{
						Name: "尖山区",
					},
					{
						Name: "岭东区",
					},
					{
						Name: "四方台区",
					},
					{
						Name: "宝山区",
					},
					{
						Name: "集贤县",
					},
					{
						Name: "友谊县",
					},
					{
						Name: "宝清县",
					},
					{
						Name: "饶河县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "大庆市",
				Children: []*Distinct{
					{
						Name: "萨尔图区",
					},
					{
						Name: "龙凤区",
					},
					{
						Name: "让胡路区",
					},
					{
						Name: "红岗区",
					},
					{
						Name: "大同区",
					},
					{
						Name: "肇州县",
					},
					{
						Name: "肇源县",
					},
					{
						Name: "林甸县",
					},
					{
						Name: "杜尔伯特蒙古族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "伊春市",
				Children: []*Distinct{
					{
						Name: "伊春区",
					},
					{
						Name: "南岔区",
					},
					{
						Name: "友好区",
					},
					{
						Name: "西林区",
					},
					{
						Name: "翠峦区",
					},
					{
						Name: "新青区",
					},
					{
						Name: "美溪区",
					},
					{
						Name: "金山屯区",
					},
					{
						Name: "五营区",
					},
					{
						Name: "乌马河区",
					},
					{
						Name: "汤旺河区",
					},
					{
						Name: "带岭区",
					},
					{
						Name: "乌伊岭区",
					},
					{
						Name: "红星区",
					},
					{
						Name: "上甘岭区",
					},
					{
						Name: "嘉荫县",
					},
					{
						Name: "铁力市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "佳木斯市",
				Children: []*Distinct{
					{
						Name: "向阳区",
					},
					{
						Name: "前进区",
					},
					{
						Name: "东风区",
					},
					{
						Name: "郊区",
					},
					{
						Name: "桦南县",
					},
					{
						Name: "桦川县",
					},
					{
						Name: "汤原县",
					},
					{
						Name: "抚远县",
					},
					{
						Name: "同江市",
					},
					{
						Name: "富锦市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "七台河市",
				Children: []*Distinct{
					{
						Name: "新兴区",
					},
					{
						Name: "桃山区",
					},
					{
						Name: "茄子河区",
					},
					{
						Name: "勃利县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "牡丹江市",
				Children: []*Distinct{
					{
						Name: "东安区",
					},
					{
						Name: "阳明区",
					},
					{
						Name: "爱民区",
					},
					{
						Name: "西安区",
					},
					{
						Name: "东宁县",
					},
					{
						Name: "林口县",
					},
					{
						Name: "绥芬河市",
					},
					{
						Name: "海林市",
					},
					{
						Name: "宁安市",
					},
					{
						Name: "穆棱市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "黑河市",
				Children: []*Distinct{
					{
						Name: "爱辉区",
					},
					{
						Name: "嫩江县",
					},
					{
						Name: "逊克县",
					},
					{
						Name: "孙吴县",
					},
					{
						Name: "北安市",
					},
					{
						Name: "五大连池市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "绥化市",
				Children: []*Distinct{
					{
						Name: "北林区",
					},
					{
						Name: "望奎县",
					},
					{
						Name: "兰西县",
					},
					{
						Name: "青冈县",
					},
					{
						Name: "庆安县",
					},
					{
						Name: "明水县",
					},
					{
						Name: "绥棱县",
					},
					{
						Name: "安达市",
					},
					{
						Name: "肇东市",
					},
					{
						Name: "海伦市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "大兴安岭地区",
				Children: []*Distinct{
					{
						Name: "松岭区",
					},
					{
						Name: "新林区",
					},
					{
						Name: "呼中区",
					},
					{
						Name: "呼玛县",
					},
					{
						Name: "塔河县",
					},
					{
						Name: "漠河县",
					},
					{
						Name: "加格达奇区",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "上海",
		Children: []*City{
			{
				Name: "上海市",
				Children: []*Distinct{
					{
						Name: "黄浦区",
					},
					{
						Name: "徐汇区",
					},
					{
						Name: "长宁区",
					},
					{
						Name: "静安区",
					},
					{
						Name: "普陀区",
					},
					{
						Name: "闸北区",
					},
					{
						Name: "虹口区",
					},
					{
						Name: "杨浦区",
					},
					{
						Name: "闵行区",
					},
					{
						Name: "宝山区",
					},
					{
						Name: "嘉定区",
					},
					{
						Name: "浦东新区",
					},
					{
						Name: "金山区",
					},
					{
						Name: "松江区",
					},
					{
						Name: "青浦区",
					},
					{
						Name: "奉贤区",
					},
				},
			},
		},
	},
	{
		Name: "江苏省",
		Children: []*City{
			{
				Name: "南京市",
				Children: []*Distinct{
					{
						Name: "玄武区",
					},
					{
						Name: "秦淮区",
					},
					{
						Name: "建邺区",
					},
					{
						Name: "鼓楼区",
					},
					{
						Name: "浦口区",
					},
					{
						Name: "栖霞区",
					},
					{
						Name: "雨花台区",
					},
					{
						Name: "江宁区",
					},
					{
						Name: "六合区",
					},
					{
						Name: "溧水区",
					},
					{
						Name: "高淳区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "无锡市",
				Children: []*Distinct{
					{
						Name: "崇安区",
					},
					{
						Name: "南长区",
					},
					{
						Name: "北塘区",
					},
					{
						Name: "锡山区",
					},
					{
						Name: "惠山区",
					},
					{
						Name: "滨湖区",
					},
					{
						Name: "江阴市",
					},
					{
						Name: "宜兴市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "徐州市",
				Children: []*Distinct{
					{
						Name: "鼓楼区",
					},
					{
						Name: "云龙区",
					},
					{
						Name: "贾汪区",
					},
					{
						Name: "泉山区",
					},
					{
						Name: "丰县",
					},
					{
						Name: "沛县",
					},
					{
						Name: "铜山区",
					},
					{
						Name: "睢宁县",
					},
					{
						Name: "新沂市",
					},
					{
						Name: "邳州市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "常州市",
				Children: []*Distinct{
					{
						Name: "天宁区",
					},
					{
						Name: "钟楼区",
					},
					{
						Name: "戚墅堰区",
					},
					{
						Name: "新北区",
					},
					{
						Name: "武进区",
					},
					{
						Name: "溧阳市",
					},
					{
						Name: "金坛市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "苏州市",
				Children: []*Distinct{
					{
						Name: "虎丘区",
					},
					{
						Name: "吴中区",
					},
					{
						Name: "相城区",
					},
					{
						Name: "姑苏区",
					},
					{
						Name: "常熟市",
					},
					{
						Name: "张家港市",
					},
					{
						Name: "昆山市",
					},
					{
						Name: "吴江区",
					},
					{
						Name: "太仓市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "南通市",
				Children: []*Distinct{
					{
						Name: "崇川区",
					},
					{
						Name: "港闸区",
					},
					{
						Name: "通州区",
					},
					{
						Name: "海安县",
					},
					{
						Name: "如东县",
					},
					{
						Name: "启东市",
					},
					{
						Name: "如皋市",
					},
					{
						Name: "海门市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "连云港市",
				Children: []*Distinct{
					{
						Name: "连云区",
					},
					{
						Name: "新浦区",
					},
					{
						Name: "海州区",
					},
					{
						Name: "赣榆县",
					},
					{
						Name: "东海县",
					},
					{
						Name: "灌云县",
					},
					{
						Name: "灌南县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "淮安市",
				Children: []*Distinct{
					{
						Name: "清河区",
					},
					{
						Name: "淮安区",
					},
					{
						Name: "淮阴区",
					},
					{
						Name: "清浦区",
					},
					{
						Name: "涟水县",
					},
					{
						Name: "洪泽县",
					},
					{
						Name: "盱眙县",
					},
					{
						Name: "金湖县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "盐城市",
				Children: []*Distinct{
					{
						Name: "亭湖区",
					},
					{
						Name: "盐都区",
					},
					{
						Name: "响水县",
					},
					{
						Name: "滨海县",
					},
					{
						Name: "阜宁县",
					},
					{
						Name: "射阳县",
					},
					{
						Name: "建湖县",
					},
					{
						Name: "东台市",
					},
					{
						Name: "大丰市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "扬州市",
				Children: []*Distinct{
					{
						Name: "广陵区",
					},
					{
						Name: "邗江区",
					},
					{
						Name: "宝应县",
					},
					{
						Name: "仪征市",
					},
					{
						Name: "高邮市",
					},
					{
						Name: "江都区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "镇江市",
				Children: []*Distinct{
					{
						Name: "京口区",
					},
					{
						Name: "润州区",
					},
					{
						Name: "丹徒区",
					},
					{
						Name: "丹阳市",
					},
					{
						Name: "扬中市",
					},
					{
						Name: "句容市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "泰州市",
				Children: []*Distinct{
					{
						Name: "海陵区",
					},
					{
						Name: "高港区",
					},
					{
						Name: "兴化市",
					},
					{
						Name: "靖江市",
					},
					{
						Name: "泰兴市",
					},
					{
						Name: "姜堰区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "宿迁市",
				Children: []*Distinct{
					{
						Name: "宿城区",
					},
					{
						Name: "宿豫区",
					},
					{
						Name: "沭阳县",
					},
					{
						Name: "泗阳县",
					},
					{
						Name: "泗洪县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "浙江省",
		Children: []*City{
			{
				Name: "杭州市",
				Children: []*Distinct{
					{
						Name: "上城区",
					},
					{
						Name: "下城区",
					},
					{
						Name: "江干区",
					},
					{
						Name: "拱墅区",
					},
					{
						Name: "西湖区",
					},
					{
						Name: "滨江区",
					},
					{
						Name: "萧山区",
					},
					{
						Name: "余杭区",
					},
					{
						Name: "桐庐县",
					},
					{
						Name: "淳安县",
					},
					{
						Name: "建德市",
					},
					{
						Name: "富阳市",
					},
					{
						Name: "临安市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "宁波市",
				Children: []*Distinct{
					{
						Name: "海曙区",
					},
					{
						Name: "江东区",
					},
					{
						Name: "江北区",
					},
					{
						Name: "北仑区",
					},
					{
						Name: "镇海区",
					},
					{
						Name: "鄞州区",
					},
					{
						Name: "象山县",
					},
					{
						Name: "宁海县",
					},
					{
						Name: "余姚市",
					},
					{
						Name: "慈溪市",
					},
					{
						Name: "奉化市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "温州市",
				Children: []*Distinct{
					{
						Name: "鹿城区",
					},
					{
						Name: "龙湾区",
					},
					{
						Name: "瓯海区",
					},
					{
						Name: "洞头县",
					},
					{
						Name: "永嘉县",
					},
					{
						Name: "平阳县",
					},
					{
						Name: "苍南县",
					},
					{
						Name: "文成县",
					},
					{
						Name: "泰顺县",
					},
					{
						Name: "瑞安市",
					},
					{
						Name: "乐清市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "嘉兴市",
				Children: []*Distinct{
					{
						Name: "南湖区",
					},
					{
						Name: "秀洲区",
					},
					{
						Name: "嘉善县",
					},
					{
						Name: "海盐县",
					},
					{
						Name: "海宁市",
					},
					{
						Name: "平湖市",
					},
					{
						Name: "桐乡市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "湖州市",
				Children: []*Distinct{
					{
						Name: "吴兴区",
					},
					{
						Name: "南浔区",
					},
					{
						Name: "德清县",
					},
					{
						Name: "长兴县",
					},
					{
						Name: "安吉县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "绍兴市",
				Children: []*Distinct{
					{
						Name: "越城区",
					},
					{
						Name: "绍兴县",
					},
					{
						Name: "新昌县",
					},
					{
						Name: "诸暨市",
					},
					{
						Name: "上虞市",
					},
					{
						Name: "嵊州市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "金华市",
				Children: []*Distinct{
					{
						Name: "婺城区",
					},
					{
						Name: "金东区",
					},
					{
						Name: "武义县",
					},
					{
						Name: "浦江县",
					},
					{
						Name: "磐安县",
					},
					{
						Name: "兰溪市",
					},
					{
						Name: "义乌市",
					},
					{
						Name: "东阳市",
					},
					{
						Name: "永康市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "衢州市",
				Children: []*Distinct{
					{
						Name: "柯城区",
					},
					{
						Name: "衢江区",
					},
					{
						Name: "常山县",
					},
					{
						Name: "开化县",
					},
					{
						Name: "龙游县",
					},
					{
						Name: "江山市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "舟山市",
				Children: []*Distinct{
					{
						Name: "定海区",
					},
					{
						Name: "普陀区",
					},
					{
						Name: "岱山县",
					},
					{
						Name: "嵊泗县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "台州市",
				Children: []*Distinct{
					{
						Name: "椒江区",
					},
					{
						Name: "黄岩区",
					},
					{
						Name: "路桥区",
					},
					{
						Name: "玉环县",
					},
					{
						Name: "三门县",
					},
					{
						Name: "天台县",
					},
					{
						Name: "仙居县",
					},
					{
						Name: "温岭市",
					},
					{
						Name: "临海市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "丽水市",
				Children: []*Distinct{
					{
						Name: "莲都区",
					},
					{
						Name: "青田县",
					},
					{
						Name: "缙云县",
					},
					{
						Name: "遂昌县",
					},
					{
						Name: "松阳县",
					},
					{
						Name: "云和县",
					},
					{
						Name: "庆元县",
					},
					{
						Name: "景宁畲族自治县",
					},
					{
						Name: "龙泉市",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "安徽省",
		Children: []*City{
			{
				Name: "合肥市",
				Children: []*Distinct{
					{
						Name: "瑶海区",
					},
					{
						Name: "庐阳区",
					},
					{
						Name: "蜀山区",
					},
					{
						Name: "包河区",
					},
					{
						Name: "长丰县",
					},
					{
						Name: "肥东县",
					},
					{
						Name: "肥西县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "芜湖市",
				Children: []*Distinct{
					{
						Name: "镜湖区",
					},
					{
						Name: "弋江区",
					},
					{
						Name: "鸠江区",
					},
					{
						Name: "三山区",
					},
					{
						Name: "芜湖县",
					},
					{
						Name: "繁昌县",
					},
					{
						Name: "南陵县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "蚌埠市",
				Children: []*Distinct{
					{
						Name: "龙子湖区",
					},
					{
						Name: "蚌山区",
					},
					{
						Name: "禹会区",
					},
					{
						Name: "淮上区",
					},
					{
						Name: "怀远县",
					},
					{
						Name: "五河县",
					},
					{
						Name: "固镇县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "淮南市",
				Children: []*Distinct{
					{
						Name: "大通区",
					},
					{
						Name: "田家庵区",
					},
					{
						Name: "谢家集区",
					},
					{
						Name: "八公山区",
					},
					{
						Name: "潘集区",
					},
					{
						Name: "凤台县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "马鞍山市",
				Children: []*Distinct{
					{
						Name: "花山区",
					},
					{
						Name: "雨山区",
					},
					{
						Name: "博望区",
					},
					{
						Name: "当涂县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "淮北市",
				Children: []*Distinct{
					{
						Name: "杜集区",
					},
					{
						Name: "相山区",
					},
					{
						Name: "烈山区",
					},
					{
						Name: "濉溪县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "铜陵市",
				Children: []*Distinct{
					{
						Name: "铜官山区",
					},
					{
						Name: "狮子山区",
					},
					{
						Name: "郊区",
					},
					{
						Name: "铜陵县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "安庆市",
				Children: []*Distinct{
					{
						Name: "迎江区",
					},
					{
						Name: "大观区",
					},
					{
						Name: "宜秀区",
					},
					{
						Name: "怀宁县",
					},
					{
						Name: "枞阳县",
					},
					{
						Name: "潜山县",
					},
					{
						Name: "太湖县",
					},
					{
						Name: "宿松县",
					},
					{
						Name: "望江县",
					},
					{
						Name: "岳西县",
					},
					{
						Name: "桐城市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "黄山市",
				Children: []*Distinct{
					{
						Name: "屯溪区",
					},
					{
						Name: "黄山区",
					},
					{
						Name: "徽州区",
					},
					{
						Name: "歙县",
					},
					{
						Name: "休宁县",
					},
					{
						Name: "黟县",
					},
					{
						Name: "祁门县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "滁州市",
				Children: []*Distinct{
					{
						Name: "琅琊区",
					},
					{
						Name: "南谯区",
					},
					{
						Name: "来安县",
					},
					{
						Name: "全椒县",
					},
					{
						Name: "定远县",
					},
					{
						Name: "凤阳县",
					},
					{
						Name: "天长市",
					},
					{
						Name: "明光市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "阜阳市",
				Children: []*Distinct{
					{
						Name: "颍州区",
					},
					{
						Name: "颍东区",
					},
					{
						Name: "颍泉区",
					},
					{
						Name: "临泉县",
					},
					{
						Name: "太和县",
					},
					{
						Name: "阜南县",
					},
					{
						Name: "颍上县",
					},
					{
						Name: "界首市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "宿州市",
				Children: []*Distinct{
					{
						Name: "埇桥区",
					},
					{
						Name: "砀山县",
					},
					{
						Name: "萧县",
					},
					{
						Name: "灵璧县",
					},
					{
						Name: "泗县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "巢湖市",
				Children: []*Distinct{
					{
						Name: "庐江县",
					},
					{
						Name: "无为县",
					},
					{
						Name: "含山县",
					},
					{
						Name: "和县",
					},
				},
			},
			{
				Name: "六安市",
				Children: []*Distinct{
					{
						Name: "金安区",
					},
					{
						Name: "裕安区",
					},
					{
						Name: "寿县",
					},
					{
						Name: "霍邱县",
					},
					{
						Name: "舒城县",
					},
					{
						Name: "金寨县",
					},
					{
						Name: "霍山县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "亳州市",
				Children: []*Distinct{
					{
						Name: "谯城区",
					},
					{
						Name: "涡阳县",
					},
					{
						Name: "蒙城县",
					},
					{
						Name: "利辛县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "池州市",
				Children: []*Distinct{
					{
						Name: "贵池区",
					},
					{
						Name: "东至县",
					},
					{
						Name: "石台县",
					},
					{
						Name: "青阳县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "宣城市",
				Children: []*Distinct{
					{
						Name: "宣州区",
					},
					{
						Name: "郎溪县",
					},
					{
						Name: "广德县",
					},
					{
						Name: "泾县",
					},
					{
						Name: "绩溪县",
					},
					{
						Name: "旌德县",
					},
					{
						Name: "宁国市",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "福建省",
		Children: []*City{
			{
				Name: "福州市",
				Children: []*Distinct{
					{
						Name: "鼓楼区",
					},
					{
						Name: "台江区",
					},
					{
						Name: "仓山区",
					},
					{
						Name: "马尾区",
					},
					{
						Name: "晋安区",
					},
					{
						Name: "闽侯县",
					},
					{
						Name: "连江县",
					},
					{
						Name: "罗源县",
					},
					{
						Name: "闽清县",
					},
					{
						Name: "永泰县",
					},
					{
						Name: "平潭县",
					},
					{
						Name: "福清市",
					},
					{
						Name: "长乐市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "厦门市",
				Children: []*Distinct{
					{
						Name: "思明区",
					},
					{
						Name: "海沧区",
					},
					{
						Name: "湖里区",
					},
					{
						Name: "集美区",
					},
					{
						Name: "同安区",
					},
					{
						Name: "翔安区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "莆田市",
				Children: []*Distinct{
					{
						Name: "城厢区",
					},
					{
						Name: "涵江区",
					},
					{
						Name: "荔城区",
					},
					{
						Name: "秀屿区",
					},
					{
						Name: "仙游县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "三明市",
				Children: []*Distinct{
					{
						Name: "梅列区",
					},
					{
						Name: "三元区",
					},
					{
						Name: "明溪县",
					},
					{
						Name: "清流县",
					},
					{
						Name: "宁化县",
					},
					{
						Name: "大田县",
					},
					{
						Name: "尤溪县",
					},
					{
						Name: "沙县",
					},
					{
						Name: "将乐县",
					},
					{
						Name: "泰宁县",
					},
					{
						Name: "建宁县",
					},
					{
						Name: "永安市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "泉州市",
				Children: []*Distinct{
					{
						Name: "鲤城区",
					},
					{
						Name: "丰泽区",
					},
					{
						Name: "洛江区",
					},
					{
						Name: "泉港区",
					},
					{
						Name: "惠安县",
					},
					{
						Name: "安溪县",
					},
					{
						Name: "永春县",
					},
					{
						Name: "德化县",
					},
					{
						Name: "金门县",
					},
					{
						Name: "石狮市",
					},
					{
						Name: "晋江市",
					},
					{
						Name: "南安市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "漳州市",
				Children: []*Distinct{
					{
						Name: "芗城区",
					},
					{
						Name: "龙文区",
					},
					{
						Name: "云霄县",
					},
					{
						Name: "漳浦县",
					},
					{
						Name: "诏安县",
					},
					{
						Name: "长泰县",
					},
					{
						Name: "东山县",
					},
					{
						Name: "南靖县",
					},
					{
						Name: "平和县",
					},
					{
						Name: "华安县",
					},
					{
						Name: "龙海市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "南平市",
				Children: []*Distinct{
					{
						Name: "延平区",
					},
					{
						Name: "顺昌县",
					},
					{
						Name: "浦城县",
					},
					{
						Name: "光泽县",
					},
					{
						Name: "松溪县",
					},
					{
						Name: "政和县",
					},
					{
						Name: "邵武市",
					},
					{
						Name: "武夷山市",
					},
					{
						Name: "建瓯市",
					},
					{
						Name: "建阳市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "龙岩市",
				Children: []*Distinct{
					{
						Name: "新罗区",
					},
					{
						Name: "长汀县",
					},
					{
						Name: "永定县",
					},
					{
						Name: "上杭县",
					},
					{
						Name: "武平县",
					},
					{
						Name: "连城县",
					},
					{
						Name: "漳平市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "宁德市",
				Children: []*Distinct{
					{
						Name: "蕉城区",
					},
					{
						Name: "霞浦县",
					},
					{
						Name: "古田县",
					},
					{
						Name: "屏南县",
					},
					{
						Name: "寿宁县",
					},
					{
						Name: "周宁县",
					},
					{
						Name: "柘荣县",
					},
					{
						Name: "福安市",
					},
					{
						Name: "福鼎市",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "江西省",
		Children: []*City{
			{
				Name: "南昌市",
				Children: []*Distinct{
					{
						Name: "东湖区",
					},
					{
						Name: "西湖区",
					},
					{
						Name: "青云谱区",
					},
					{
						Name: "湾里区",
					},
					{
						Name: "青山湖区",
					},
					{
						Name: "南昌县",
					},
					{
						Name: "新建县",
					},
					{
						Name: "安义县",
					},
					{
						Name: "进贤县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "景德镇市",
				Children: []*Distinct{
					{
						Name: "昌江区",
					},
					{
						Name: "珠山区",
					},
					{
						Name: "浮梁县",
					},
					{
						Name: "乐平市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "萍乡市",
				Children: []*Distinct{
					{
						Name: "安源区",
					},
					{
						Name: "湘东区",
					},
					{
						Name: "莲花县",
					},
					{
						Name: "上栗县",
					},
					{
						Name: "芦溪县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "九江市",
				Children: []*Distinct{
					{
						Name: "庐山区",
					},
					{
						Name: "浔阳区",
					},
					{
						Name: "九江县",
					},
					{
						Name: "武宁县",
					},
					{
						Name: "修水县",
					},
					{
						Name: "永修县",
					},
					{
						Name: "德安县",
					},
					{
						Name: "星子县",
					},
					{
						Name: "都昌县",
					},
					{
						Name: "湖口县",
					},
					{
						Name: "彭泽县",
					},
					{
						Name: "瑞昌市",
					},
					{
						Name: "其它区",
					},
					{
						Name: "共青城市",
					},
				},
			},
			{
				Name: "新余市",
				Children: []*Distinct{
					{
						Name: "渝水区",
					},
					{
						Name: "分宜县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "鹰潭市",
				Children: []*Distinct{
					{
						Name: "月湖区",
					},
					{
						Name: "余江县",
					},
					{
						Name: "贵溪市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "赣州市",
				Children: []*Distinct{
					{
						Name: "章贡区",
					},
					{
						Name: "赣县",
					},
					{
						Name: "信丰县",
					},
					{
						Name: "大余县",
					},
					{
						Name: "上犹县",
					},
					{
						Name: "崇义县",
					},
					{
						Name: "安远县",
					},
					{
						Name: "龙南县",
					},
					{
						Name: "定南县",
					},
					{
						Name: "全南县",
					},
					{
						Name: "宁都县",
					},
					{
						Name: "于都县",
					},
					{
						Name: "兴国县",
					},
					{
						Name: "会昌县",
					},
					{
						Name: "寻乌县",
					},
					{
						Name: "石城县",
					},
					{
						Name: "瑞金市",
					},
					{
						Name: "南康市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "吉安市",
				Children: []*Distinct{
					{
						Name: "吉州区",
					},
					{
						Name: "青原区",
					},
					{
						Name: "吉安县",
					},
					{
						Name: "吉水县",
					},
					{
						Name: "峡江县",
					},
					{
						Name: "新干县",
					},
					{
						Name: "永丰县",
					},
					{
						Name: "泰和县",
					},
					{
						Name: "遂川县",
					},
					{
						Name: "万安县",
					},
					{
						Name: "安福县",
					},
					{
						Name: "永新县",
					},
					{
						Name: "井冈山市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "宜春市",
				Children: []*Distinct{
					{
						Name: "袁州区",
					},
					{
						Name: "奉新县",
					},
					{
						Name: "万载县",
					},
					{
						Name: "上高县",
					},
					{
						Name: "宜丰县",
					},
					{
						Name: "靖安县",
					},
					{
						Name: "铜鼓县",
					},
					{
						Name: "丰城市",
					},
					{
						Name: "樟树市",
					},
					{
						Name: "高安市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "抚州市",
				Children: []*Distinct{
					{
						Name: "临川区",
					},
					{
						Name: "南城县",
					},
					{
						Name: "黎川县",
					},
					{
						Name: "南丰县",
					},
					{
						Name: "崇仁县",
					},
					{
						Name: "乐安县",
					},
					{
						Name: "宜黄县",
					},
					{
						Name: "金溪县",
					},
					{
						Name: "资溪县",
					},
					{
						Name: "东乡县",
					},
					{
						Name: "广昌县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "上饶市",
				Children: []*Distinct{
					{
						Name: "信州区",
					},
					{
						Name: "上饶县",
					},
					{
						Name: "广丰县",
					},
					{
						Name: "玉山县",
					},
					{
						Name: "铅山县",
					},
					{
						Name: "横峰县",
					},
					{
						Name: "弋阳县",
					},
					{
						Name: "余干县",
					},
					{
						Name: "鄱阳县",
					},
					{
						Name: "万年县",
					},
					{
						Name: "婺源县",
					},
					{
						Name: "德兴市",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "山东省",
		Children: []*City{
			{
				Name: "济南市",
				Children: []*Distinct{
					{
						Name: "历下区",
					},
					{
						Name: "市中区",
					},
					{
						Name: "槐荫区",
					},
					{
						Name: "天桥区",
					},
					{
						Name: "历城区",
					},
					{
						Name: "长清区",
					},
					{
						Name: "平阴县",
					},
					{
						Name: "济阳县",
					},
					{
						Name: "商河县",
					},
					{
						Name: "章丘市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "青岛市",
				Children: []*Distinct{
					{
						Name: "市南区",
					},
					{
						Name: "市北区",
					},
					{
						Name: "黄岛区",
					},
					{
						Name: "崂山区",
					},
					{
						Name: "李沧区",
					},
					{
						Name: "城阳区",
					},
					{
						Name: "胶州市",
					},
					{
						Name: "即墨市",
					},
					{
						Name: "平度市",
					},
					{
						Name: "莱西市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "淄博市",
				Children: []*Distinct{
					{
						Name: "淄川区",
					},
					{
						Name: "张店区",
					},
					{
						Name: "博山区",
					},
					{
						Name: "临淄区",
					},
					{
						Name: "周村区",
					},
					{
						Name: "桓台县",
					},
					{
						Name: "高青县",
					},
					{
						Name: "沂源县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "枣庄市",
				Children: []*Distinct{
					{
						Name: "市中区",
					},
					{
						Name: "薛城区",
					},
					{
						Name: "峄城区",
					},
					{
						Name: "台儿庄区",
					},
					{
						Name: "山亭区",
					},
					{
						Name: "滕州市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "东营市",
				Children: []*Distinct{
					{
						Name: "东营区",
					},
					{
						Name: "河口区",
					},
					{
						Name: "垦利县",
					},
					{
						Name: "利津县",
					},
					{
						Name: "广饶县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "烟台市",
				Children: []*Distinct{
					{
						Name: "芝罘区",
					},
					{
						Name: "福山区",
					},
					{
						Name: "牟平区",
					},
					{
						Name: "莱山区",
					},
					{
						Name: "长岛县",
					},
					{
						Name: "龙口市",
					},
					{
						Name: "莱阳市",
					},
					{
						Name: "莱州市",
					},
					{
						Name: "蓬莱市",
					},
					{
						Name: "招远市",
					},
					{
						Name: "栖霞市",
					},
					{
						Name: "海阳市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "潍坊市",
				Children: []*Distinct{
					{
						Name: "潍城区",
					},
					{
						Name: "寒亭区",
					},
					{
						Name: "坊子区",
					},
					{
						Name: "奎文区",
					},
					{
						Name: "临朐县",
					},
					{
						Name: "昌乐县",
					},
					{
						Name: "青州市",
					},
					{
						Name: "诸城市",
					},
					{
						Name: "寿光市",
					},
					{
						Name: "安丘市",
					},
					{
						Name: "高密市",
					},
					{
						Name: "昌邑市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "济宁市",
				Children: []*Distinct{
					{
						Name: "市中区",
					},
					{
						Name: "任城区",
					},
					{
						Name: "微山县",
					},
					{
						Name: "鱼台县",
					},
					{
						Name: "金乡县",
					},
					{
						Name: "嘉祥县",
					},
					{
						Name: "汶上县",
					},
					{
						Name: "泗水县",
					},
					{
						Name: "梁山县",
					},
					{
						Name: "曲阜市",
					},
					{
						Name: "兖州市",
					},
					{
						Name: "邹城市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "泰安市",
				Children: []*Distinct{
					{
						Name: "泰山区",
					},
					{
						Name: "岱岳区",
					},
					{
						Name: "宁阳县",
					},
					{
						Name: "东平县",
					},
					{
						Name: "新泰市",
					},
					{
						Name: "肥城市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "威海市",
				Children: []*Distinct{
					{
						Name: "环翠区",
					},
					{
						Name: "文登市",
					},
					{
						Name: "荣成市",
					},
					{
						Name: "乳山市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "日照市",
				Children: []*Distinct{
					{
						Name: "东港区",
					},
					{
						Name: "岚山区",
					},
					{
						Name: "五莲县",
					},
					{
						Name: "莒县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "莱芜市",
				Children: []*Distinct{
					{
						Name: "莱城区",
					},
					{
						Name: "钢城区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "临沂市",
				Children: []*Distinct{
					{
						Name: "兰山区",
					},
					{
						Name: "罗庄区",
					},
					{
						Name: "河东区",
					},
					{
						Name: "沂南县",
					},
					{
						Name: "郯城县",
					},
					{
						Name: "沂水县",
					},
					{
						Name: "苍山县",
					},
					{
						Name: "费县",
					},
					{
						Name: "平邑县",
					},
					{
						Name: "莒南县",
					},
					{
						Name: "蒙阴县",
					},
					{
						Name: "临沭县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "德州市",
				Children: []*Distinct{
					{
						Name: "德城区",
					},
					{
						Name: "陵县",
					},
					{
						Name: "宁津县",
					},
					{
						Name: "庆云县",
					},
					{
						Name: "临邑县",
					},
					{
						Name: "齐河县",
					},
					{
						Name: "平原县",
					},
					{
						Name: "夏津县",
					},
					{
						Name: "武城县",
					},
					{
						Name: "乐陵市",
					},
					{
						Name: "禹城市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "聊城市",
				Children: []*Distinct{
					{
						Name: "东昌府区",
					},
					{
						Name: "阳谷县",
					},
					{
						Name: "莘县",
					},
					{
						Name: "茌平县",
					},
					{
						Name: "东阿县",
					},
					{
						Name: "冠县",
					},
					{
						Name: "高唐县",
					},
					{
						Name: "临清市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "滨州市",
				Children: []*Distinct{
					{
						Name: "滨城区",
					},
					{
						Name: "惠民县",
					},
					{
						Name: "阳信县",
					},
					{
						Name: "无棣县",
					},
					{
						Name: "沾化县",
					},
					{
						Name: "博兴县",
					},
					{
						Name: "邹平县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "菏泽市",
				Children: []*Distinct{
					{
						Name: "牡丹区",
					},
					{
						Name: "曹县",
					},
					{
						Name: "单县",
					},
					{
						Name: "成武县",
					},
					{
						Name: "巨野县",
					},
					{
						Name: "郓城县",
					},
					{
						Name: "鄄城县",
					},
					{
						Name: "定陶县",
					},
					{
						Name: "东明县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "河南省",
		Children: []*City{
			{
				Name: "郑州市",
				Children: []*Distinct{
					{
						Name: "中原区",
					},
					{
						Name: "二七区",
					},
					{
						Name: "管城回族区",
					},
					{
						Name: "金水区",
					},
					{
						Name: "上街区",
					},
					{
						Name: "惠济区",
					},
					{
						Name: "中牟县",
					},
					{
						Name: "巩义市",
					},
					{
						Name: "荥阳市",
					},
					{
						Name: "新密市",
					},
					{
						Name: "新郑市",
					},
					{
						Name: "登封市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "开封市",
				Children: []*Distinct{
					{
						Name: "龙亭区",
					},
					{
						Name: "顺河回族区",
					},
					{
						Name: "鼓楼区",
					},
					{
						Name: "禹王台区",
					},
					{
						Name: "金明区",
					},
					{
						Name: "杞县",
					},
					{
						Name: "通许县",
					},
					{
						Name: "尉氏县",
					},
					{
						Name: "开封县",
					},
					{
						Name: "兰考县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "洛阳市",
				Children: []*Distinct{
					{
						Name: "老城区",
					},
					{
						Name: "西工区",
					},
					{
						Name: "瀍河回族区",
					},
					{
						Name: "涧西区",
					},
					{
						Name: "吉利区",
					},
					{
						Name: "洛龙区",
					},
					{
						Name: "孟津县",
					},
					{
						Name: "新安县",
					},
					{
						Name: "栾川县",
					},
					{
						Name: "嵩县",
					},
					{
						Name: "汝阳县",
					},
					{
						Name: "宜阳县",
					},
					{
						Name: "洛宁县",
					},
					{
						Name: "伊川县",
					},
					{
						Name: "偃师市",
					},
				},
			},
			{
				Name: "平顶山市",
				Children: []*Distinct{
					{
						Name: "新华区",
					},
					{
						Name: "卫东区",
					},
					{
						Name: "石龙区",
					},
					{
						Name: "湛河区",
					},
					{
						Name: "宝丰县",
					},
					{
						Name: "叶县",
					},
					{
						Name: "鲁山县",
					},
					{
						Name: "郏县",
					},
					{
						Name: "舞钢市",
					},
					{
						Name: "汝州市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "安阳市",
				Children: []*Distinct{
					{
						Name: "文峰区",
					},
					{
						Name: "北关区",
					},
					{
						Name: "殷都区",
					},
					{
						Name: "龙安区",
					},
					{
						Name: "安阳县",
					},
					{
						Name: "汤阴县",
					},
					{
						Name: "滑县",
					},
					{
						Name: "内黄县",
					},
					{
						Name: "林州市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "鹤壁市",
				Children: []*Distinct{
					{
						Name: "鹤山区",
					},
					{
						Name: "山城区",
					},
					{
						Name: "淇滨区",
					},
					{
						Name: "浚县",
					},
					{
						Name: "淇县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "新乡市",
				Children: []*Distinct{
					{
						Name: "红旗区",
					},
					{
						Name: "卫滨区",
					},
					{
						Name: "凤泉区",
					},
					{
						Name: "牧野区",
					},
					{
						Name: "新乡县",
					},
					{
						Name: "获嘉县",
					},
					{
						Name: "原阳县",
					},
					{
						Name: "延津县",
					},
					{
						Name: "封丘县",
					},
					{
						Name: "长垣县",
					},
					{
						Name: "卫辉市",
					},
					{
						Name: "辉县市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "焦作市",
				Children: []*Distinct{
					{
						Name: "解放区",
					},
					{
						Name: "中站区",
					},
					{
						Name: "马村区",
					},
					{
						Name: "山阳区",
					},
					{
						Name: "修武县",
					},
					{
						Name: "博爱县",
					},
					{
						Name: "武陟县",
					},
					{
						Name: "温县",
					},
					{
						Name: "济源市",
					},
					{
						Name: "沁阳市",
					},
					{
						Name: "孟州市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "濮阳市",
				Children: []*Distinct{
					{
						Name: "华龙区",
					},
					{
						Name: "清丰县",
					},
					{
						Name: "南乐县",
					},
					{
						Name: "范县",
					},
					{
						Name: "台前县",
					},
					{
						Name: "濮阳县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "许昌市",
				Children: []*Distinct{
					{
						Name: "魏都区",
					},
					{
						Name: "许昌县",
					},
					{
						Name: "鄢陵县",
					},
					{
						Name: "襄城县",
					},
					{
						Name: "禹州市",
					},
					{
						Name: "长葛市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "漯河市",
				Children: []*Distinct{
					{
						Name: "源汇区",
					},
					{
						Name: "郾城区",
					},
					{
						Name: "召陵区",
					},
					{
						Name: "舞阳县",
					},
					{
						Name: "临颍县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "三门峡市",
				Children: []*Distinct{
					{
						Name: "湖滨区",
					},
					{
						Name: "渑池县",
					},
					{
						Name: "陕县",
					},
					{
						Name: "卢氏县",
					},
					{
						Name: "义马市",
					},
					{
						Name: "灵宝市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "南阳市",
				Children: []*Distinct{
					{
						Name: "宛城区",
					},
					{
						Name: "卧龙区",
					},
					{
						Name: "南召县",
					},
					{
						Name: "方城县",
					},
					{
						Name: "西峡县",
					},
					{
						Name: "镇平县",
					},
					{
						Name: "内乡县",
					},
					{
						Name: "淅川县",
					},
					{
						Name: "社旗县",
					},
					{
						Name: "唐河县",
					},
					{
						Name: "新野县",
					},
					{
						Name: "桐柏县",
					},
					{
						Name: "邓州市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "商丘市",
				Children: []*Distinct{
					{
						Name: "梁园区",
					},
					{
						Name: "睢阳区",
					},
					{
						Name: "民权县",
					},
					{
						Name: "睢县",
					},
					{
						Name: "宁陵县",
					},
					{
						Name: "柘城县",
					},
					{
						Name: "虞城县",
					},
					{
						Name: "夏邑县",
					},
					{
						Name: "永城市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "信阳市",
				Children: []*Distinct{
					{
						Name: "浉河区",
					},
					{
						Name: "平桥区",
					},
					{
						Name: "罗山县",
					},
					{
						Name: "光山县",
					},
					{
						Name: "新县",
					},
					{
						Name: "商城县",
					},
					{
						Name: "固始县",
					},
					{
						Name: "潢川县",
					},
					{
						Name: "淮滨县",
					},
					{
						Name: "息县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "周口市",
				Children: []*Distinct{
					{
						Name: "川汇区",
					},
					{
						Name: "扶沟县",
					},
					{
						Name: "西华县",
					},
					{
						Name: "商水县",
					},
					{
						Name: "沈丘县",
					},
					{
						Name: "郸城县",
					},
					{
						Name: "淮阳县",
					},
					{
						Name: "太康县",
					},
					{
						Name: "鹿邑县",
					},
					{
						Name: "项城市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "驻马店市",
				Children: []*Distinct{
					{
						Name: "驿城区",
					},
					{
						Name: "西平县",
					},
					{
						Name: "上蔡县",
					},
					{
						Name: "平舆县",
					},
					{
						Name: "正阳县",
					},
					{
						Name: "确山县",
					},
					{
						Name: "泌阳县",
					},
					{
						Name: "汝南县",
					},
					{
						Name: "遂平县",
					},
					{
						Name: "新蔡县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "湖北省",
		Children: []*City{
			{
				Name: "武汉市",
				Children: []*Distinct{
					{
						Name: "江岸区",
					},
					{
						Name: "江汉区",
					},
					{
						Name: "硚口区",
					},
					{
						Name: "汉阳区",
					},
					{
						Name: "武昌区",
					},
					{
						Name: "青山区",
					},
					{
						Name: "洪山区",
					},
					{
						Name: "东西湖区",
					},
					{
						Name: "汉南区",
					},
					{
						Name: "蔡甸区",
					},
					{
						Name: "江夏区",
					},
					{
						Name: "黄陂区",
					},
					{
						Name: "新洲区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "黄石市",
				Children: []*Distinct{
					{
						Name: "黄石港区",
					},
					{
						Name: "西塞山区",
					},
					{
						Name: "下陆区",
					},
					{
						Name: "铁山区",
					},
					{
						Name: "阳新县",
					},
					{
						Name: "大冶市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "十堰市",
				Children: []*Distinct{
					{
						Name: "茅箭区",
					},
					{
						Name: "张湾区",
					},
					{
						Name: "郧县",
					},
					{
						Name: "郧西县",
					},
					{
						Name: "竹山县",
					},
					{
						Name: "竹溪县",
					},
					{
						Name: "房县",
					},
					{
						Name: "丹江口市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "宜昌市",
				Children: []*Distinct{
					{
						Name: "西陵区",
					},
					{
						Name: "伍家岗区",
					},
					{
						Name: "点军区",
					},
					{
						Name: "猇亭区",
					},
					{
						Name: "夷陵区",
					},
					{
						Name: "远安县",
					},
					{
						Name: "兴山县",
					},
					{
						Name: "秭归县",
					},
					{
						Name: "长阳土家族自治县",
					},
					{
						Name: "五峰土家族自治县",
					},
					{
						Name: "宜都市",
					},
					{
						Name: "当阳市",
					},
					{
						Name: "枝江市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "襄阳市",
				Children: []*Distinct{
					{
						Name: "襄城区",
					},
					{
						Name: "樊城区",
					},
					{
						Name: "襄州区",
					},
					{
						Name: "南漳县",
					},
					{
						Name: "谷城县",
					},
					{
						Name: "保康县",
					},
					{
						Name: "老河口市",
					},
					{
						Name: "枣阳市",
					},
					{
						Name: "宜城市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "鄂州市",
				Children: []*Distinct{
					{
						Name: "梁子湖区",
					},
					{
						Name: "华容区",
					},
					{
						Name: "鄂城区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "荆门市",
				Children: []*Distinct{
					{
						Name: "东宝区",
					},
					{
						Name: "掇刀区",
					},
					{
						Name: "京山县",
					},
					{
						Name: "沙洋县",
					},
					{
						Name: "钟祥市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "孝感市",
				Children: []*Distinct{
					{
						Name: "孝南区",
					},
					{
						Name: "孝昌县",
					},
					{
						Name: "大悟县",
					},
					{
						Name: "云梦县",
					},
					{
						Name: "应城市",
					},
					{
						Name: "安陆市",
					},
					{
						Name: "汉川市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "荆州市",
				Children: []*Distinct{
					{
						Name: "沙市区",
					},
					{
						Name: "荆州区",
					},
					{
						Name: "公安县",
					},
					{
						Name: "监利县",
					},
					{
						Name: "江陵县",
					},
					{
						Name: "石首市",
					},
					{
						Name: "洪湖市",
					},
					{
						Name: "松滋市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "黄冈市",
				Children: []*Distinct{
					{
						Name: "黄州区",
					},
					{
						Name: "团风县",
					},
					{
						Name: "红安县",
					},
					{
						Name: "罗田县",
					},
					{
						Name: "英山县",
					},
					{
						Name: "浠水县",
					},
					{
						Name: "蕲春县",
					},
					{
						Name: "黄梅县",
					},
					{
						Name: "麻城市",
					},
					{
						Name: "武穴市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "咸宁市",
				Children: []*Distinct{
					{
						Name: "咸安区",
					},
					{
						Name: "嘉鱼县",
					},
					{
						Name: "通城县",
					},
					{
						Name: "崇阳县",
					},
					{
						Name: "通山县",
					},
					{
						Name: "赤壁市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "随州市",
				Children: []*Distinct{
					{
						Name: "曾都区",
					},
					{
						Name: "随县",
					},
					{
						Name: "广水市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "恩施土家族苗族自治州",
				Children: []*Distinct{
					{
						Name: "恩施市",
					},
					{
						Name: "利川市",
					},
					{
						Name: "建始县",
					},
					{
						Name: "巴东县",
					},
					{
						Name: "宣恩县",
					},
					{
						Name: "咸丰县",
					},
					{
						Name: "来凤县",
					},
					{
						Name: "鹤峰县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "湖南省",
		Children: []*City{
			{
				Name: "长沙市",
				Children: []*Distinct{
					{
						Name: "芙蓉区",
					},
					{
						Name: "天心区",
					},
					{
						Name: "岳麓区",
					},
					{
						Name: "开福区",
					},
					{
						Name: "雨花区",
					},
					{
						Name: "长沙县",
					},
					{
						Name: "望城区",
					},
					{
						Name: "宁乡县",
					},
					{
						Name: "浏阳市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "株洲市",
				Children: []*Distinct{
					{
						Name: "荷塘区",
					},
					{
						Name: "芦淞区",
					},
					{
						Name: "石峰区",
					},
					{
						Name: "天元区",
					},
					{
						Name: "株洲县",
					},
					{
						Name: "攸县",
					},
					{
						Name: "茶陵县",
					},
					{
						Name: "炎陵县",
					},
					{
						Name: "醴陵市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "湘潭市",
				Children: []*Distinct{
					{
						Name: "雨湖区",
					},
					{
						Name: "岳塘区",
					},
					{
						Name: "湘潭县",
					},
					{
						Name: "湘乡市",
					},
					{
						Name: "韶山市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "衡阳市",
				Children: []*Distinct{
					{
						Name: "珠晖区",
					},
					{
						Name: "雁峰区",
					},
					{
						Name: "石鼓区",
					},
					{
						Name: "蒸湘区",
					},
					{
						Name: "南岳区",
					},
					{
						Name: "衡阳县",
					},
					{
						Name: "衡南县",
					},
					{
						Name: "衡山县",
					},
					{
						Name: "衡东县",
					},
					{
						Name: "祁东县",
					},
					{
						Name: "耒阳市",
					},
					{
						Name: "常宁市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "邵阳市",
				Children: []*Distinct{
					{
						Name: "双清区",
					},
					{
						Name: "大祥区",
					},
					{
						Name: "北塔区",
					},
					{
						Name: "邵东县",
					},
					{
						Name: "新邵县",
					},
					{
						Name: "邵阳县",
					},
					{
						Name: "隆回县",
					},
					{
						Name: "洞口县",
					},
					{
						Name: "绥宁县",
					},
					{
						Name: "新宁县",
					},
					{
						Name: "城步苗族自治县",
					},
					{
						Name: "武冈市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "岳阳市",
				Children: []*Distinct{
					{
						Name: "岳阳楼区",
					},
					{
						Name: "云溪区",
					},
					{
						Name: "君山区",
					},
					{
						Name: "岳阳县",
					},
					{
						Name: "华容县",
					},
					{
						Name: "湘阴县",
					},
					{
						Name: "平江县",
					},
					{
						Name: "汨罗市",
					},
					{
						Name: "临湘市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "常德市",
				Children: []*Distinct{
					{
						Name: "武陵区",
					},
					{
						Name: "鼎城区",
					},
					{
						Name: "安乡县",
					},
					{
						Name: "汉寿县",
					},
					{
						Name: "澧县",
					},
					{
						Name: "临澧县",
					},
					{
						Name: "桃源县",
					},
					{
						Name: "石门县",
					},
					{
						Name: "津市市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "张家界市",
				Children: []*Distinct{
					{
						Name: "永定区",
					},
					{
						Name: "武陵源区",
					},
					{
						Name: "慈利县",
					},
					{
						Name: "桑植县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "益阳市",
				Children: []*Distinct{
					{
						Name: "资阳区",
					},
					{
						Name: "赫山区",
					},
					{
						Name: "南县",
					},
					{
						Name: "桃江县",
					},
					{
						Name: "安化县",
					},
					{
						Name: "沅江市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "郴州市",
				Children: []*Distinct{
					{
						Name: "北湖区",
					},
					{
						Name: "苏仙区",
					},
					{
						Name: "桂阳县",
					},
					{
						Name: "宜章县",
					},
					{
						Name: "永兴县",
					},
					{
						Name: "嘉禾县",
					},
					{
						Name: "临武县",
					},
					{
						Name: "汝城县",
					},
					{
						Name: "桂东县",
					},
					{
						Name: "安仁县",
					},
					{
						Name: "资兴市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "永州市",
				Children: []*Distinct{
					{
						Name: "零陵区",
					},
					{
						Name: "冷水滩区",
					},
					{
						Name: "祁阳县",
					},
					{
						Name: "东安县",
					},
					{
						Name: "双牌县",
					},
					{
						Name: "道县",
					},
					{
						Name: "江永县",
					},
					{
						Name: "宁远县",
					},
					{
						Name: "蓝山县",
					},
					{
						Name: "新田县",
					},
					{
						Name: "江华瑶族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "怀化市",
				Children: []*Distinct{
					{
						Name: "鹤城区",
					},
					{
						Name: "中方县",
					},
					{
						Name: "沅陵县",
					},
					{
						Name: "辰溪县",
					},
					{
						Name: "溆浦县",
					},
					{
						Name: "会同县",
					},
					{
						Name: "麻阳苗族自治县",
					},
					{
						Name: "新晃侗族自治县",
					},
					{
						Name: "芷江侗族自治县",
					},
					{
						Name: "靖州苗族侗族自治县",
					},
					{
						Name: "通道侗族自治县",
					},
					{
						Name: "洪江市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "娄底市",
				Children: []*Distinct{
					{
						Name: "娄星区",
					},
					{
						Name: "双峰县",
					},
					{
						Name: "新化县",
					},
					{
						Name: "冷水江市",
					},
					{
						Name: "涟源市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "湘西土家族苗族自治州",
				Children: []*Distinct{
					{
						Name: "吉首市",
					},
					{
						Name: "泸溪县",
					},
					{
						Name: "凤凰县",
					},
					{
						Name: "花垣县",
					},
					{
						Name: "保靖县",
					},
					{
						Name: "古丈县",
					},
					{
						Name: "永顺县",
					},
					{
						Name: "龙山县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "广东省",
		Children: []*City{
			{
				Name: "广州市",
				Children: []*Distinct{
					{
						Name: "荔湾区",
					},
					{
						Name: "越秀区",
					},
					{
						Name: "海珠区",
					},
					{
						Name: "天河区",
					},
					{
						Name: "白云区",
					},
					{
						Name: "黄埔区",
					},
					{
						Name: "番禺区",
					},
					{
						Name: "花都区",
					},
					{
						Name: "南沙区",
					},
					{
						Name: "萝岗区",
					},
					{
						Name: "增城市",
					},
					{
						Name: "从化市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "韶关市",
				Children: []*Distinct{
					{
						Name: "武江区",
					},
					{
						Name: "浈江区",
					},
					{
						Name: "曲江区",
					},
					{
						Name: "始兴县",
					},
					{
						Name: "仁化县",
					},
					{
						Name: "翁源县",
					},
					{
						Name: "乳源瑶族自治县",
					},
					{
						Name: "新丰县",
					},
					{
						Name: "乐昌市",
					},
					{
						Name: "南雄市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "深圳市",
				Children: []*Distinct{
					{
						Name: "罗湖区",
					},
					{
						Name: "福田区",
					},
					{
						Name: "南山区",
					},
					{
						Name: "宝安区",
					},
					{
						Name: "龙岗区",
					},
					{
						Name: "盐田区",
					},
					{
						Name: "其它区",
					},
					{
						Name: "光明新区",
					},
					{
						Name: "坪山新区",
					},
					{
						Name: "大鹏新区",
					},
					{
						Name: "龙华新区",
					},
				},
			},
			{
				Name: "珠海市",
				Children: []*Distinct{
					{
						Name: "香洲区",
					},
					{
						Name: "斗门区",
					},
					{
						Name: "金湾区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "汕头市",
				Children: []*Distinct{
					{
						Name: "龙湖区",
					},
					{
						Name: "金平区",
					},
					{
						Name: "濠江区",
					},
					{
						Name: "潮阳区",
					},
					{
						Name: "潮南区",
					},
					{
						Name: "澄海区",
					},
					{
						Name: "南澳县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "佛山市",
				Children: []*Distinct{
					{
						Name: "禅城区",
					},
					{
						Name: "南海区",
					},
					{
						Name: "顺德区",
					},
					{
						Name: "三水区",
					},
					{
						Name: "高明区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "江门市",
				Children: []*Distinct{
					{
						Name: "蓬江区",
					},
					{
						Name: "江海区",
					},
					{
						Name: "新会区",
					},
					{
						Name: "台山市",
					},
					{
						Name: "开平市",
					},
					{
						Name: "鹤山市",
					},
					{
						Name: "恩平市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "湛江市",
				Children: []*Distinct{
					{
						Name: "赤坎区",
					},
					{
						Name: "霞山区",
					},
					{
						Name: "坡头区",
					},
					{
						Name: "麻章区",
					},
					{
						Name: "遂溪县",
					},
					{
						Name: "徐闻县",
					},
					{
						Name: "廉江市",
					},
					{
						Name: "雷州市",
					},
					{
						Name: "吴川市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "茂名市",
				Children: []*Distinct{
					{
						Name: "茂南区",
					},
					{
						Name: "茂港区",
					},
					{
						Name: "电白县",
					},
					{
						Name: "高州市",
					},
					{
						Name: "化州市",
					},
					{
						Name: "信宜市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "肇庆市",
				Children: []*Distinct{
					{
						Name: "端州区",
					},
					{
						Name: "鼎湖区",
					},
					{
						Name: "广宁县",
					},
					{
						Name: "怀集县",
					},
					{
						Name: "封开县",
					},
					{
						Name: "德庆县",
					},
					{
						Name: "高要市",
					},
					{
						Name: "四会市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "惠州市",
				Children: []*Distinct{
					{
						Name: "惠城区",
					},
					{
						Name: "惠阳区",
					},
					{
						Name: "博罗县",
					},
					{
						Name: "惠东县",
					},
					{
						Name: "龙门县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "梅州市",
				Children: []*Distinct{
					{
						Name: "梅江区",
					},
					{
						Name: "梅县",
					},
					{
						Name: "大埔县",
					},
					{
						Name: "丰顺县",
					},
					{
						Name: "五华县",
					},
					{
						Name: "平远县",
					},
					{
						Name: "蕉岭县",
					},
					{
						Name: "兴宁市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "汕尾市",
				Children: []*Distinct{
					{
						Name: "城区",
					},
					{
						Name: "海丰县",
					},
					{
						Name: "陆河县",
					},
					{
						Name: "陆丰市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "河源市",
				Children: []*Distinct{
					{
						Name: "源城区",
					},
					{
						Name: "紫金县",
					},
					{
						Name: "龙川县",
					},
					{
						Name: "连平县",
					},
					{
						Name: "和平县",
					},
					{
						Name: "东源县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "阳江市",
				Children: []*Distinct{
					{
						Name: "江城区",
					},
					{
						Name: "阳西县",
					},
					{
						Name: "阳东县",
					},
					{
						Name: "阳春市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "清远市",
				Children: []*Distinct{
					{
						Name: "清城区",
					},
					{
						Name: "佛冈县",
					},
					{
						Name: "阳山县",
					},
					{
						Name: "连山壮族瑶族自治县",
					},
					{
						Name: "连南瑶族自治县",
					},
					{
						Name: "清新区",
					},
					{
						Name: "英德市",
					},
					{
						Name: "连州市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "东莞市",
			},
			{
				Name: "中山市",
			},
			{
				Name: "潮州市",
				Children: []*Distinct{
					{
						Name: "湘桥区",
					},
					{
						Name: "潮安区",
					},
					{
						Name: "饶平县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "揭阳市",
				Children: []*Distinct{
					{
						Name: "榕城区",
					},
					{
						Name: "揭东区",
					},
					{
						Name: "揭西县",
					},
					{
						Name: "惠来县",
					},
					{
						Name: "普宁市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "云浮市",
				Children: []*Distinct{
					{
						Name: "云城区",
					},
					{
						Name: "新兴县",
					},
					{
						Name: "郁南县",
					},
					{
						Name: "云安县",
					},
					{
						Name: "罗定市",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "广西壮族自治区",
		Children: []*City{
			{
				Name: "南宁市",
				Children: []*Distinct{
					{
						Name: "兴宁区",
					},
					{
						Name: "青秀区",
					},
					{
						Name: "江南区",
					},
					{
						Name: "西乡塘区",
					},
					{
						Name: "良庆区",
					},
					{
						Name: "邕宁区",
					},
					{
						Name: "武鸣县",
					},
					{
						Name: "隆安县",
					},
					{
						Name: "马山县",
					},
					{
						Name: "上林县",
					},
					{
						Name: "宾阳县",
					},
					{
						Name: "横县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "柳州市",
				Children: []*Distinct{
					{
						Name: "城中区",
					},
					{
						Name: "鱼峰区",
					},
					{
						Name: "柳南区",
					},
					{
						Name: "柳北区",
					},
					{
						Name: "柳江县",
					},
					{
						Name: "柳城县",
					},
					{
						Name: "鹿寨县",
					},
					{
						Name: "融安县",
					},
					{
						Name: "融水苗族自治县",
					},
					{
						Name: "三江侗族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "桂林市",
				Children: []*Distinct{
					{
						Name: "秀峰区",
					},
					{
						Name: "叠彩区",
					},
					{
						Name: "象山区",
					},
					{
						Name: "七星区",
					},
					{
						Name: "雁山区",
					},
					{
						Name: "阳朔县",
					},
					{
						Name: "临桂区",
					},
					{
						Name: "灵川县",
					},
					{
						Name: "全州县",
					},
					{
						Name: "兴安县",
					},
					{
						Name: "永福县",
					},
					{
						Name: "灌阳县",
					},
					{
						Name: "龙胜各族自治县",
					},
					{
						Name: "资源县",
					},
					{
						Name: "平乐县",
					},
					{
						Name: "荔浦县",
					},
					{
						Name: "恭城瑶族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "梧州市",
				Children: []*Distinct{
					{
						Name: "万秀区",
					},
					{
						Name: "长洲区",
					},
					{
						Name: "龙圩区",
					},
					{
						Name: "苍梧县",
					},
					{
						Name: "藤县",
					},
					{
						Name: "蒙山县",
					},
					{
						Name: "岑溪市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "北海市",
				Children: []*Distinct{
					{
						Name: "海城区",
					},
					{
						Name: "银海区",
					},
					{
						Name: "铁山港区",
					},
					{
						Name: "合浦县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "防城港市",
				Children: []*Distinct{
					{
						Name: "港口区",
					},
					{
						Name: "防城区",
					},
					{
						Name: "上思县",
					},
					{
						Name: "东兴市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "钦州市",
				Children: []*Distinct{
					{
						Name: "钦南区",
					},
					{
						Name: "钦北区",
					},
					{
						Name: "灵山县",
					},
					{
						Name: "浦北县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "贵港市",
				Children: []*Distinct{
					{
						Name: "港北区",
					},
					{
						Name: "港南区",
					},
					{
						Name: "覃塘区",
					},
					{
						Name: "平南县",
					},
					{
						Name: "桂平市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "玉林市",
				Children: []*Distinct{
					{
						Name: "玉州区",
					},
					{
						Name: "福绵区",
					},
					{
						Name: "容县",
					},
					{
						Name: "陆川县",
					},
					{
						Name: "博白县",
					},
					{
						Name: "兴业县",
					},
					{
						Name: "北流市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "百色市",
				Children: []*Distinct{
					{
						Name: "右江区",
					},
					{
						Name: "田阳县",
					},
					{
						Name: "田东县",
					},
					{
						Name: "平果县",
					},
					{
						Name: "德保县",
					},
					{
						Name: "靖西县",
					},
					{
						Name: "那坡县",
					},
					{
						Name: "凌云县",
					},
					{
						Name: "乐业县",
					},
					{
						Name: "田林县",
					},
					{
						Name: "西林县",
					},
					{
						Name: "隆林各族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "贺州市",
				Children: []*Distinct{
					{
						Name: "八步区",
					},
					{
						Name: "平桂管理区",
					},
					{
						Name: "昭平县",
					},
					{
						Name: "钟山县",
					},
					{
						Name: "富川瑶族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "河池市",
				Children: []*Distinct{
					{
						Name: "金城江区",
					},
					{
						Name: "南丹县",
					},
					{
						Name: "天峨县",
					},
					{
						Name: "凤山县",
					},
					{
						Name: "东兰县",
					},
					{
						Name: "罗城仫佬族自治县",
					},
					{
						Name: "环江毛南族自治县",
					},
					{
						Name: "巴马瑶族自治县",
					},
					{
						Name: "都安瑶族自治县",
					},
					{
						Name: "大化瑶族自治县",
					},
					{
						Name: "宜州市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "来宾市",
				Children: []*Distinct{
					{
						Name: "兴宾区",
					},
					{
						Name: "忻城县",
					},
					{
						Name: "象州县",
					},
					{
						Name: "武宣县",
					},
					{
						Name: "金秀瑶族自治县",
					},
					{
						Name: "合山市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "崇左市",
				Children: []*Distinct{
					{
						Name: "江州区",
					},
					{
						Name: "扶绥县",
					},
					{
						Name: "宁明县",
					},
					{
						Name: "龙州县",
					},
					{
						Name: "大新县",
					},
					{
						Name: "天等县",
					},
					{
						Name: "凭祥市",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "海南省",
		Children: []*City{
			{
				Name: "海口市",
				Children: []*Distinct{
					{
						Name: "秀英区",
					},
					{
						Name: "龙华区",
					},
					{
						Name: "琼山区",
					},
					{
						Name: "美兰区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "三亚市",
			},
			{
				Name: "三沙市",
				Children: []*Distinct{
					{
						Name: "西沙群岛",
					},
					{
						Name: "南沙群岛",
					},
					{
						Name: "中沙群岛的岛礁及其海域",
					},
				},
			},
		},
	},
	{
		Name: "重庆",
		Children: []*City{
			{
				Name: "重庆市",
				Children: []*Distinct{
					{
						Name: "万州区",
					},
					{
						Name: "涪陵区",
					},
					{
						Name: "渝中区",
					},
					{
						Name: "大渡口区",
					},
					{
						Name: "江北区",
					},
					{
						Name: "沙坪坝区",
					},
					{
						Name: "九龙坡区",
					},
					{
						Name: "南岸区",
					},
					{
						Name: "北碚区",
					},
					{
						Name: "万盛区",
					},
					{
						Name: "双桥区",
					},
					{
						Name: "渝北区",
					},
					{
						Name: "巴南区",
					},
					{
						Name: "黔江区",
					},
					{
						Name: "长寿区",
					},
				},
			},
		},
	},
	{
		Name: "四川省",
		Children: []*City{
			{
				Name: "成都市",
				Children: []*Distinct{
					{
						Name: "锦江区",
					},
					{
						Name: "青羊区",
					},
					{
						Name: "金牛区",
					},
					{
						Name: "武侯区",
					},
					{
						Name: "成华区",
					},
					{
						Name: "龙泉驿区",
					},
					{
						Name: "青白江区",
					},
					{
						Name: "新都区",
					},
					{
						Name: "温江区",
					},
					{
						Name: "金堂县",
					},
					{
						Name: "双流县",
					},
					{
						Name: "郫县",
					},
					{
						Name: "大邑县",
					},
					{
						Name: "蒲江县",
					},
					{
						Name: "新津县",
					},
					{
						Name: "都江堰市",
					},
					{
						Name: "彭州市",
					},
					{
						Name: "邛崃市",
					},
					{
						Name: "崇州市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "自贡市",
				Children: []*Distinct{
					{
						Name: "自流井区",
					},
					{
						Name: "贡井区",
					},
					{
						Name: "大安区",
					},
					{
						Name: "沿滩区",
					},
					{
						Name: "荣县",
					},
					{
						Name: "富顺县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "攀枝花市",
				Children: []*Distinct{
					{
						Name: "东区",
					},
					{
						Name: "西区",
					},
					{
						Name: "仁和区",
					},
					{
						Name: "米易县",
					},
					{
						Name: "盐边县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "泸州市",
				Children: []*Distinct{
					{
						Name: "江阳区",
					},
					{
						Name: "纳溪区",
					},
					{
						Name: "龙马潭区",
					},
					{
						Name: "泸县",
					},
					{
						Name: "合江县",
					},
					{
						Name: "叙永县",
					},
					{
						Name: "古蔺县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "德阳市",
				Children: []*Distinct{
					{
						Name: "旌阳区",
					},
					{
						Name: "中江县",
					},
					{
						Name: "罗江县",
					},
					{
						Name: "广汉市",
					},
					{
						Name: "什邡市",
					},
					{
						Name: "绵竹市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "绵阳市",
				Children: []*Distinct{
					{
						Name: "涪城区",
					},
					{
						Name: "游仙区",
					},
					{
						Name: "三台县",
					},
					{
						Name: "盐亭县",
					},
					{
						Name: "安县",
					},
					{
						Name: "梓潼县",
					},
					{
						Name: "北川羌族自治县",
					},
					{
						Name: "平武县",
					},
					{
						Name: "江油市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "广元市",
				Children: []*Distinct{
					{
						Name: "利州区",
					},
					{
						Name: "昭化区",
					},
					{
						Name: "朝天区",
					},
					{
						Name: "旺苍县",
					},
					{
						Name: "青川县",
					},
					{
						Name: "剑阁县",
					},
					{
						Name: "苍溪县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "遂宁市",
				Children: []*Distinct{
					{
						Name: "船山区",
					},
					{
						Name: "安居区",
					},
					{
						Name: "蓬溪县",
					},
					{
						Name: "射洪县",
					},
					{
						Name: "大英县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "内江市",
				Children: []*Distinct{
					{
						Name: "市中区",
					},
					{
						Name: "东兴区",
					},
					{
						Name: "威远县",
					},
					{
						Name: "资中县",
					},
					{
						Name: "隆昌县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "乐山市",
				Children: []*Distinct{
					{
						Name: "市中区",
					},
					{
						Name: "沙湾区",
					},
					{
						Name: "五通桥区",
					},
					{
						Name: "金口河区",
					},
					{
						Name: "犍为县",
					},
					{
						Name: "井研县",
					},
					{
						Name: "夹江县",
					},
					{
						Name: "沐川县",
					},
					{
						Name: "峨边彝族自治县",
					},
					{
						Name: "马边彝族自治县",
					},
					{
						Name: "峨眉山市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "南充市",
				Children: []*Distinct{
					{
						Name: "顺庆区",
					},
					{
						Name: "高坪区",
					},
					{
						Name: "嘉陵区",
					},
					{
						Name: "南部县",
					},
					{
						Name: "营山县",
					},
					{
						Name: "蓬安县",
					},
					{
						Name: "仪陇县",
					},
					{
						Name: "西充县",
					},
					{
						Name: "阆中市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "眉山市",
				Children: []*Distinct{
					{
						Name: "东坡区",
					},
					{
						Name: "仁寿县",
					},
					{
						Name: "彭山县",
					},
					{
						Name: "洪雅县",
					},
					{
						Name: "丹棱县",
					},
					{
						Name: "青神县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "宜宾市",
				Children: []*Distinct{
					{
						Name: "翠屏区",
					},
					{
						Name: "宜宾县",
					},
					{
						Name: "南溪区",
					},
					{
						Name: "江安县",
					},
					{
						Name: "长宁县",
					},
					{
						Name: "高县",
					},
					{
						Name: "珙县",
					},
					{
						Name: "筠连县",
					},
					{
						Name: "兴文县",
					},
					{
						Name: "屏山县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "广安市",
				Children: []*Distinct{
					{
						Name: "广安区",
					},
					{
						Name: "前锋区",
					},
					{
						Name: "岳池县",
					},
					{
						Name: "武胜县",
					},
					{
						Name: "邻水县",
					},
					{
						Name: "华蓥市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "达州市",
				Children: []*Distinct{
					{
						Name: "通川区",
					},
					{
						Name: "达川区",
					},
					{
						Name: "宣汉县",
					},
					{
						Name: "开江县",
					},
					{
						Name: "大竹县",
					},
					{
						Name: "渠县",
					},
					{
						Name: "万源市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "雅安市",
				Children: []*Distinct{
					{
						Name: "雨城区",
					},
					{
						Name: "名山区",
					},
					{
						Name: "荥经县",
					},
					{
						Name: "汉源县",
					},
					{
						Name: "石棉县",
					},
					{
						Name: "天全县",
					},
					{
						Name: "芦山县",
					},
					{
						Name: "宝兴县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "巴中市",
				Children: []*Distinct{
					{
						Name: "巴州区",
					},
					{
						Name: "恩阳区",
					},
					{
						Name: "通江县",
					},
					{
						Name: "南江县",
					},
					{
						Name: "平昌县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "资阳市",
				Children: []*Distinct{
					{
						Name: "雁江区",
					},
					{
						Name: "安岳县",
					},
					{
						Name: "乐至县",
					},
					{
						Name: "简阳市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "阿坝藏族羌族自治州",
				Children: []*Distinct{
					{
						Name: "汶川县",
					},
					{
						Name: "理县",
					},
					{
						Name: "茂县",
					},
					{
						Name: "松潘县",
					},
					{
						Name: "九寨沟县",
					},
					{
						Name: "金川县",
					},
					{
						Name: "小金县",
					},
					{
						Name: "黑水县",
					},
					{
						Name: "马尔康县",
					},
					{
						Name: "壤塘县",
					},
					{
						Name: "阿坝县",
					},
					{
						Name: "若尔盖县",
					},
					{
						Name: "红原县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "甘孜藏族自治州",
				Children: []*Distinct{
					{
						Name: "康定县",
					},
					{
						Name: "泸定县",
					},
					{
						Name: "丹巴县",
					},
					{
						Name: "九龙县",
					},
					{
						Name: "雅江县",
					},
					{
						Name: "道孚县",
					},
					{
						Name: "炉霍县",
					},
					{
						Name: "甘孜县",
					},
					{
						Name: "新龙县",
					},
					{
						Name: "德格县",
					},
					{
						Name: "白玉县",
					},
					{
						Name: "石渠县",
					},
					{
						Name: "色达县",
					},
					{
						Name: "理塘县",
					},
					{
						Name: "巴塘县",
					},
					{
						Name: "乡城县",
					},
					{
						Name: "稻城县",
					},
					{
						Name: "得荣县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "凉山彝族自治州",
				Children: []*Distinct{
					{
						Name: "西昌市",
					},
					{
						Name: "木里藏族自治县",
					},
					{
						Name: "盐源县",
					},
					{
						Name: "德昌县",
					},
					{
						Name: "会理县",
					},
					{
						Name: "会东县",
					},
					{
						Name: "宁南县",
					},
					{
						Name: "普格县",
					},
					{
						Name: "布拖县",
					},
					{
						Name: "金阳县",
					},
					{
						Name: "昭觉县",
					},
					{
						Name: "喜德县",
					},
					{
						Name: "冕宁县",
					},
					{
						Name: "越西县",
					},
					{
						Name: "甘洛县",
					},
					{
						Name: "美姑县",
					},
					{
						Name: "雷波县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "贵州省",
		Children: []*City{
			{
				Name: "贵阳市",
				Children: []*Distinct{
					{
						Name: "南明区",
					},
					{
						Name: "云岩区",
					},
					{
						Name: "花溪区",
					},
					{
						Name: "乌当区",
					},
					{
						Name: "白云区",
					},
					{
						Name: "开阳县",
					},
					{
						Name: "息烽县",
					},
					{
						Name: "修文县",
					},
					{
						Name: "观山湖区",
					},
					{
						Name: "清镇市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "六盘水市",
				Children: []*Distinct{
					{
						Name: "钟山区",
					},
					{
						Name: "六枝特区",
					},
					{
						Name: "水城县",
					},
					{
						Name: "盘县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "遵义市",
				Children: []*Distinct{
					{
						Name: "红花岗区",
					},
					{
						Name: "汇川区",
					},
					{
						Name: "遵义县",
					},
					{
						Name: "桐梓县",
					},
					{
						Name: "绥阳县",
					},
					{
						Name: "正安县",
					},
					{
						Name: "道真仡佬族苗族自治县",
					},
					{
						Name: "务川仡佬族苗族自治县",
					},
					{
						Name: "凤冈县",
					},
					{
						Name: "湄潭县",
					},
					{
						Name: "余庆县",
					},
					{
						Name: "习水县",
					},
					{
						Name: "赤水市",
					},
					{
						Name: "仁怀市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "安顺市",
				Children: []*Distinct{
					{
						Name: "西秀区",
					},
					{
						Name: "平坝县",
					},
					{
						Name: "普定县",
					},
					{
						Name: "镇宁布依族苗族自治县",
					},
					{
						Name: "关岭布依族苗族自治县",
					},
					{
						Name: "紫云苗族布依族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "铜仁市",
				Children: []*Distinct{
					{
						Name: "碧江区",
					},
					{
						Name: "江口县",
					},
					{
						Name: "玉屏侗族自治县",
					},
					{
						Name: "石阡县",
					},
					{
						Name: "思南县",
					},
					{
						Name: "印江土家族苗族自治县",
					},
					{
						Name: "德江县",
					},
					{
						Name: "沿河土家族自治县",
					},
					{
						Name: "松桃苗族自治县",
					},
					{
						Name: "万山区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "黔西南布依族苗族自治州",
				Children: []*Distinct{
					{
						Name: "兴义市",
					},
					{
						Name: "兴仁县",
					},
					{
						Name: "普安县",
					},
					{
						Name: "晴隆县",
					},
					{
						Name: "贞丰县",
					},
					{
						Name: "望谟县",
					},
					{
						Name: "册亨县",
					},
					{
						Name: "安龙县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "毕节市",
				Children: []*Distinct{
					{
						Name: "七星关区",
					},
					{
						Name: "大方县",
					},
					{
						Name: "黔西县",
					},
					{
						Name: "金沙县",
					},
					{
						Name: "织金县",
					},
					{
						Name: "纳雍县",
					},
					{
						Name: "威宁彝族回族苗族自治县",
					},
					{
						Name: "赫章县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "黔东南苗族侗族自治州",
				Children: []*Distinct{
					{
						Name: "凯里市",
					},
					{
						Name: "黄平县",
					},
					{
						Name: "施秉县",
					},
					{
						Name: "三穗县",
					},
					{
						Name: "镇远县",
					},
					{
						Name: "岑巩县",
					},
					{
						Name: "天柱县",
					},
					{
						Name: "锦屏县",
					},
					{
						Name: "剑河县",
					},
					{
						Name: "台江县",
					},
					{
						Name: "黎平县",
					},
					{
						Name: "榕江县",
					},
					{
						Name: "从江县",
					},
					{
						Name: "雷山县",
					},
					{
						Name: "麻江县",
					},
					{
						Name: "丹寨县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "黔南布依族苗族自治州",
				Children: []*Distinct{
					{
						Name: "都匀市",
					},
					{
						Name: "福泉市",
					},
					{
						Name: "荔波县",
					},
					{
						Name: "贵定县",
					},
					{
						Name: "瓮安县",
					},
					{
						Name: "独山县",
					},
					{
						Name: "平塘县",
					},
					{
						Name: "罗甸县",
					},
					{
						Name: "长顺县",
					},
					{
						Name: "龙里县",
					},
					{
						Name: "惠水县",
					},
					{
						Name: "三都水族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "云南省",
		Children: []*City{
			{
				Name: "昆明市",
				Children: []*Distinct{
					{
						Name: "五华区",
					},
					{
						Name: "盘龙区",
					},
					{
						Name: "官渡区",
					},
					{
						Name: "西山区",
					},
					{
						Name: "东川区",
					},
					{
						Name: "呈贡区",
					},
					{
						Name: "晋宁县",
					},
					{
						Name: "富民县",
					},
					{
						Name: "宜良县",
					},
					{
						Name: "石林彝族自治县",
					},
					{
						Name: "嵩明县",
					},
					{
						Name: "禄劝彝族苗族自治县",
					},
					{
						Name: "寻甸回族彝族自治县",
					},
					{
						Name: "安宁市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "曲靖市",
				Children: []*Distinct{
					{
						Name: "麒麟区",
					},
					{
						Name: "马龙县",
					},
					{
						Name: "陆良县",
					},
					{
						Name: "师宗县",
					},
					{
						Name: "罗平县",
					},
					{
						Name: "富源县",
					},
					{
						Name: "会泽县",
					},
					{
						Name: "沾益县",
					},
					{
						Name: "宣威市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "玉溪市",
				Children: []*Distinct{
					{
						Name: "红塔区",
					},
					{
						Name: "江川县",
					},
					{
						Name: "澄江县",
					},
					{
						Name: "通海县",
					},
					{
						Name: "华宁县",
					},
					{
						Name: "易门县",
					},
					{
						Name: "峨山彝族自治县",
					},
					{
						Name: "新平彝族傣族自治县",
					},
					{
						Name: "元江哈尼族彝族傣族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "保山市",
				Children: []*Distinct{
					{
						Name: "隆阳区",
					},
					{
						Name: "施甸县",
					},
					{
						Name: "腾冲县",
					},
					{
						Name: "龙陵县",
					},
					{
						Name: "昌宁县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "昭通市",
				Children: []*Distinct{
					{
						Name: "昭阳区",
					},
					{
						Name: "鲁甸县",
					},
					{
						Name: "巧家县",
					},
					{
						Name: "盐津县",
					},
					{
						Name: "大关县",
					},
					{
						Name: "永善县",
					},
					{
						Name: "绥江县",
					},
					{
						Name: "镇雄县",
					},
					{
						Name: "彝良县",
					},
					{
						Name: "威信县",
					},
					{
						Name: "水富县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "丽江市",
				Children: []*Distinct{
					{
						Name: "古城区",
					},
					{
						Name: "玉龙纳西族自治县",
					},
					{
						Name: "永胜县",
					},
					{
						Name: "华坪县",
					},
					{
						Name: "宁蒗彝族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "普洱市",
				Children: []*Distinct{
					{
						Name: "思茅区",
					},
					{
						Name: "宁洱哈尼族彝族自治县",
					},
					{
						Name: "墨江哈尼族自治县",
					},
					{
						Name: "景东彝族自治县",
					},
					{
						Name: "景谷傣族彝族自治县",
					},
					{
						Name: "镇沅彝族哈尼族拉祜族自治县",
					},
					{
						Name: "江城哈尼族彝族自治县",
					},
					{
						Name: "孟连傣族拉祜族佤族自治县",
					},
					{
						Name: "澜沧拉祜族自治县",
					},
					{
						Name: "西盟佤族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "临沧市",
				Children: []*Distinct{
					{
						Name: "临翔区",
					},
					{
						Name: "凤庆县",
					},
					{
						Name: "云县",
					},
					{
						Name: "永德县",
					},
					{
						Name: "镇康县",
					},
					{
						Name: "双江拉祜族佤族布朗族傣族自治县",
					},
					{
						Name: "耿马傣族佤族自治县",
					},
					{
						Name: "沧源佤族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "楚雄彝族自治州",
				Children: []*Distinct{
					{
						Name: "楚雄市",
					},
					{
						Name: "双柏县",
					},
					{
						Name: "牟定县",
					},
					{
						Name: "南华县",
					},
					{
						Name: "姚安县",
					},
					{
						Name: "大姚县",
					},
					{
						Name: "永仁县",
					},
					{
						Name: "元谋县",
					},
					{
						Name: "武定县",
					},
					{
						Name: "禄丰县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "红河哈尼族彝族自治州",
				Children: []*Distinct{
					{
						Name: "个旧市",
					},
					{
						Name: "开远市",
					},
					{
						Name: "蒙自市",
					},
					{
						Name: "屏边苗族自治县",
					},
					{
						Name: "建水县",
					},
					{
						Name: "石屏县",
					},
					{
						Name: "弥勒市",
					},
					{
						Name: "泸西县",
					},
					{
						Name: "元阳县",
					},
					{
						Name: "红河县",
					},
					{
						Name: "金平苗族瑶族傣族自治县",
					},
					{
						Name: "绿春县",
					},
					{
						Name: "河口瑶族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "文山壮族苗族自治州",
				Children: []*Distinct{
					{
						Name: "文山市",
					},
					{
						Name: "砚山县",
					},
					{
						Name: "西畴县",
					},
					{
						Name: "麻栗坡县",
					},
					{
						Name: "马关县",
					},
					{
						Name: "丘北县",
					},
					{
						Name: "广南县",
					},
					{
						Name: "富宁县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "西双版纳傣族自治州",
				Children: []*Distinct{
					{
						Name: "景洪市",
					},
					{
						Name: "勐海县",
					},
					{
						Name: "勐腊县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "大理白族自治州",
				Children: []*Distinct{
					{
						Name: "大理市",
					},
					{
						Name: "漾濞彝族自治县",
					},
					{
						Name: "祥云县",
					},
					{
						Name: "宾川县",
					},
					{
						Name: "弥渡县",
					},
					{
						Name: "南涧彝族自治县",
					},
					{
						Name: "巍山彝族回族自治县",
					},
					{
						Name: "永平县",
					},
					{
						Name: "云龙县",
					},
					{
						Name: "洱源县",
					},
					{
						Name: "剑川县",
					},
					{
						Name: "鹤庆县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "德宏傣族景颇族自治州",
				Children: []*Distinct{
					{
						Name: "瑞丽市",
					},
					{
						Name: "芒市",
					},
					{
						Name: "梁河县",
					},
					{
						Name: "盈江县",
					},
					{
						Name: "陇川县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "怒江傈僳族自治州",
				Children: []*Distinct{
					{
						Name: "泸水县",
					},
					{
						Name: "福贡县",
					},
					{
						Name: "贡山独龙族怒族自治县",
					},
					{
						Name: "兰坪白族普米族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "迪庆藏族自治州",
				Children: []*Distinct{
					{
						Name: "香格里拉县",
					},
					{
						Name: "德钦县",
					},
					{
						Name: "维西傈僳族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "西藏自治区",
		Children: []*City{
			{
				Name: "拉萨市",
				Children: []*Distinct{
					{
						Name: "城关区",
					},
					{
						Name: "林周县",
					},
					{
						Name: "当雄县",
					},
					{
						Name: "尼木县",
					},
					{
						Name: "曲水县",
					},
					{
						Name: "堆龙德庆县",
					},
					{
						Name: "达孜县",
					},
					{
						Name: "墨竹工卡县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "昌都地区",
				Children: []*Distinct{
					{
						Name: "昌都县",
					},
					{
						Name: "江达县",
					},
					{
						Name: "贡觉县",
					},
					{
						Name: "类乌齐县",
					},
					{
						Name: "丁青县",
					},
					{
						Name: "察雅县",
					},
					{
						Name: "八宿县",
					},
					{
						Name: "左贡县",
					},
					{
						Name: "芒康县",
					},
					{
						Name: "洛隆县",
					},
					{
						Name: "边坝县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "山南地区",
				Children: []*Distinct{
					{
						Name: "乃东县",
					},
					{
						Name: "扎囊县",
					},
					{
						Name: "贡嘎县",
					},
					{
						Name: "桑日县",
					},
					{
						Name: "琼结县",
					},
					{
						Name: "曲松县",
					},
					{
						Name: "措美县",
					},
					{
						Name: "洛扎县",
					},
					{
						Name: "加查县",
					},
					{
						Name: "隆子县",
					},
					{
						Name: "错那县",
					},
					{
						Name: "浪卡子县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "日喀则地区",
				Children: []*Distinct{
					{
						Name: "日喀则市",
					},
					{
						Name: "南木林县",
					},
					{
						Name: "江孜县",
					},
					{
						Name: "定日县",
					},
					{
						Name: "萨迦县",
					},
					{
						Name: "拉孜县",
					},
					{
						Name: "昂仁县",
					},
					{
						Name: "谢通门县",
					},
					{
						Name: "白朗县",
					},
					{
						Name: "仁布县",
					},
					{
						Name: "康马县",
					},
					{
						Name: "定结县",
					},
					{
						Name: "仲巴县",
					},
					{
						Name: "亚东县",
					},
					{
						Name: "吉隆县",
					},
					{
						Name: "聂拉木县",
					},
					{
						Name: "萨嘎县",
					},
					{
						Name: "岗巴县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "那曲地区",
				Children: []*Distinct{
					{
						Name: "那曲县",
					},
					{
						Name: "嘉黎县",
					},
					{
						Name: "比如县",
					},
					{
						Name: "聂荣县",
					},
					{
						Name: "安多县",
					},
					{
						Name: "申扎县",
					},
					{
						Name: "索县",
					},
					{
						Name: "班戈县",
					},
					{
						Name: "巴青县",
					},
					{
						Name: "尼玛县",
					},
					{
						Name: "其它区",
					},
					{
						Name: "双湖县",
					},
				},
			},
			{
				Name: "阿里地区",
				Children: []*Distinct{
					{
						Name: "普兰县",
					},
					{
						Name: "札达县",
					},
					{
						Name: "噶尔县",
					},
					{
						Name: "日土县",
					},
					{
						Name: "革吉县",
					},
					{
						Name: "改则县",
					},
					{
						Name: "措勤县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "林芝地区",
				Children: []*Distinct{
					{
						Name: "林芝县",
					},
					{
						Name: "工布江达县",
					},
					{
						Name: "米林县",
					},
					{
						Name: "墨脱县",
					},
					{
						Name: "波密县",
					},
					{
						Name: "察隅县",
					},
					{
						Name: "朗县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "陕西省",
		Children: []*City{
			{
				Name: "西安市",
				Children: []*Distinct{
					{
						Name: "新城区",
					},
					{
						Name: "碑林区",
					},
					{
						Name: "莲湖区",
					},
					{
						Name: "灞桥区",
					},
					{
						Name: "未央区",
					},
					{
						Name: "雁塔区",
					},
					{
						Name: "阎良区",
					},
					{
						Name: "临潼区",
					},
					{
						Name: "长安区",
					},
					{
						Name: "蓝田县",
					},
					{
						Name: "周至县",
					},
					{
						Name: "户县",
					},
					{
						Name: "高陵县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "铜川市",
				Children: []*Distinct{
					{
						Name: "王益区",
					},
					{
						Name: "印台区",
					},
					{
						Name: "耀州区",
					},
					{
						Name: "宜君县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "宝鸡市",
				Children: []*Distinct{
					{
						Name: "渭滨区",
					},
					{
						Name: "金台区",
					},
					{
						Name: "陈仓区",
					},
					{
						Name: "凤翔县",
					},
					{
						Name: "岐山县",
					},
					{
						Name: "扶风县",
					},
					{
						Name: "眉县",
					},
					{
						Name: "陇县",
					},
					{
						Name: "千阳县",
					},
					{
						Name: "麟游县",
					},
					{
						Name: "凤县",
					},
					{
						Name: "太白县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "咸阳市",
				Children: []*Distinct{
					{
						Name: "秦都区",
					},
					{
						Name: "杨陵区",
					},
					{
						Name: "渭城区",
					},
					{
						Name: "三原县",
					},
					{
						Name: "泾阳县",
					},
					{
						Name: "乾县",
					},
					{
						Name: "礼泉县",
					},
					{
						Name: "永寿县",
					},
					{
						Name: "彬县",
					},
					{
						Name: "长武县",
					},
					{
						Name: "旬邑县",
					},
					{
						Name: "淳化县",
					},
					{
						Name: "武功县",
					},
					{
						Name: "兴平市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "渭南市",
				Children: []*Distinct{
					{
						Name: "临渭区",
					},
					{
						Name: "华县",
					},
					{
						Name: "潼关县",
					},
					{
						Name: "大荔县",
					},
					{
						Name: "合阳县",
					},
					{
						Name: "澄城县",
					},
					{
						Name: "蒲城县",
					},
					{
						Name: "白水县",
					},
					{
						Name: "富平县",
					},
					{
						Name: "韩城市",
					},
					{
						Name: "华阴市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "延安市",
				Children: []*Distinct{
					{
						Name: "宝塔区",
					},
					{
						Name: "延长县",
					},
					{
						Name: "延川县",
					},
					{
						Name: "子长县",
					},
					{
						Name: "安塞县",
					},
					{
						Name: "志丹县",
					},
					{
						Name: "吴起县",
					},
					{
						Name: "甘泉县",
					},
					{
						Name: "富县",
					},
					{
						Name: "洛川县",
					},
					{
						Name: "宜川县",
					},
					{
						Name: "黄龙县",
					},
					{
						Name: "黄陵县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "汉中市",
				Children: []*Distinct{
					{
						Name: "汉台区",
					},
					{
						Name: "南郑县",
					},
					{
						Name: "城固县",
					},
					{
						Name: "洋县",
					},
					{
						Name: "西乡县",
					},
					{
						Name: "勉县",
					},
					{
						Name: "宁强县",
					},
					{
						Name: "略阳县",
					},
					{
						Name: "镇巴县",
					},
					{
						Name: "留坝县",
					},
					{
						Name: "佛坪县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "榆林市",
				Children: []*Distinct{
					{
						Name: "榆阳区",
					},
					{
						Name: "神木县",
					},
					{
						Name: "府谷县",
					},
					{
						Name: "横山县",
					},
					{
						Name: "靖边县",
					},
					{
						Name: "定边县",
					},
					{
						Name: "绥德县",
					},
					{
						Name: "米脂县",
					},
					{
						Name: "佳县",
					},
					{
						Name: "吴堡县",
					},
					{
						Name: "清涧县",
					},
					{
						Name: "子洲县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "安康市",
				Children: []*Distinct{
					{
						Name: "汉滨区",
					},
					{
						Name: "汉阴县",
					},
					{
						Name: "石泉县",
					},
					{
						Name: "宁陕县",
					},
					{
						Name: "紫阳县",
					},
					{
						Name: "岚皋县",
					},
					{
						Name: "平利县",
					},
					{
						Name: "镇坪县",
					},
					{
						Name: "旬阳县",
					},
					{
						Name: "白河县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "商洛市",
				Children: []*Distinct{
					{
						Name: "商州区",
					},
					{
						Name: "洛南县",
					},
					{
						Name: "丹凤县",
					},
					{
						Name: "商南县",
					},
					{
						Name: "山阳县",
					},
					{
						Name: "镇安县",
					},
					{
						Name: "柞水县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "甘肃省",
		Children: []*City{
			{
				Name: "兰州市",
				Children: []*Distinct{
					{
						Name: "城关区",
					},
					{
						Name: "七里河区",
					},
					{
						Name: "西固区",
					},
					{
						Name: "安宁区",
					},
					{
						Name: "红古区",
					},
					{
						Name: "永登县",
					},
					{
						Name: "皋兰县",
					},
					{
						Name: "榆中县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "嘉峪关市",
			},
			{
				Name: "金昌市",
				Children: []*Distinct{
					{
						Name: "金川区",
					},
					{
						Name: "永昌县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "白银市",
				Children: []*Distinct{
					{
						Name: "白银区",
					},
					{
						Name: "平川区",
					},
					{
						Name: "靖远县",
					},
					{
						Name: "会宁县",
					},
					{
						Name: "景泰县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "天水市",
				Children: []*Distinct{
					{
						Name: "秦州区",
					},
					{
						Name: "麦积区",
					},
					{
						Name: "清水县",
					},
					{
						Name: "秦安县",
					},
					{
						Name: "甘谷县",
					},
					{
						Name: "武山县",
					},
					{
						Name: "张家川回族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "武威市",
				Children: []*Distinct{
					{
						Name: "凉州区",
					},
					{
						Name: "民勤县",
					},
					{
						Name: "古浪县",
					},
					{
						Name: "天祝藏族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "张掖市",
				Children: []*Distinct{
					{
						Name: "甘州区",
					},
					{
						Name: "肃南裕固族自治县",
					},
					{
						Name: "民乐县",
					},
					{
						Name: "临泽县",
					},
					{
						Name: "高台县",
					},
					{
						Name: "山丹县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "平凉市",
				Children: []*Distinct{
					{
						Name: "崆峒区",
					},
					{
						Name: "泾川县",
					},
					{
						Name: "灵台县",
					},
					{
						Name: "崇信县",
					},
					{
						Name: "华亭县",
					},
					{
						Name: "庄浪县",
					},
					{
						Name: "静宁县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "酒泉市",
				Children: []*Distinct{
					{
						Name: "肃州区",
					},
					{
						Name: "金塔县",
					},
					{
						Name: "瓜州县",
					},
					{
						Name: "肃北蒙古族自治县",
					},
					{
						Name: "阿克塞哈萨克族自治县",
					},
					{
						Name: "玉门市",
					},
					{
						Name: "敦煌市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "庆阳市",
				Children: []*Distinct{
					{
						Name: "西峰区",
					},
					{
						Name: "庆城县",
					},
					{
						Name: "环县",
					},
					{
						Name: "华池县",
					},
					{
						Name: "合水县",
					},
					{
						Name: "正宁县",
					},
					{
						Name: "宁县",
					},
					{
						Name: "镇原县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "定西市",
				Children: []*Distinct{
					{
						Name: "安定区",
					},
					{
						Name: "通渭县",
					},
					{
						Name: "陇西县",
					},
					{
						Name: "渭源县",
					},
					{
						Name: "临洮县",
					},
					{
						Name: "漳县",
					},
					{
						Name: "岷县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "陇南市",
				Children: []*Distinct{
					{
						Name: "武都区",
					},
					{
						Name: "成县",
					},
					{
						Name: "文县",
					},
					{
						Name: "宕昌县",
					},
					{
						Name: "康县",
					},
					{
						Name: "西和县",
					},
					{
						Name: "礼县",
					},
					{
						Name: "徽县",
					},
					{
						Name: "两当县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "临夏回族自治州",
				Children: []*Distinct{
					{
						Name: "临夏市",
					},
					{
						Name: "临夏县",
					},
					{
						Name: "康乐县",
					},
					{
						Name: "永靖县",
					},
					{
						Name: "广河县",
					},
					{
						Name: "和政县",
					},
					{
						Name: "东乡族自治县",
					},
					{
						Name: "积石山保安族东乡族撒拉族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "甘南藏族自治州",
				Children: []*Distinct{
					{
						Name: "合作市",
					},
					{
						Name: "临潭县",
					},
					{
						Name: "卓尼县",
					},
					{
						Name: "舟曲县",
					},
					{
						Name: "迭部县",
					},
					{
						Name: "玛曲县",
					},
					{
						Name: "碌曲县",
					},
					{
						Name: "夏河县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "青海省",
		Children: []*City{
			{
				Name: "西宁市",
				Children: []*Distinct{
					{
						Name: "城东区",
					},
					{
						Name: "城中区",
					},
					{
						Name: "城西区",
					},
					{
						Name: "城北区",
					},
					{
						Name: "大通回族土族自治县",
					},
					{
						Name: "湟中县",
					},
					{
						Name: "湟源县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "海东市",
				Children: []*Distinct{
					{
						Name: "平安县",
					},
					{
						Name: "民和回族土族自治县",
					},
					{
						Name: "乐都区",
					},
					{
						Name: "互助土族自治县",
					},
					{
						Name: "化隆回族自治县",
					},
					{
						Name: "循化撒拉族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "海北藏族自治州",
				Children: []*Distinct{
					{
						Name: "门源回族自治县",
					},
					{
						Name: "祁连县",
					},
					{
						Name: "海晏县",
					},
					{
						Name: "刚察县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "黄南藏族自治州",
				Children: []*Distinct{
					{
						Name: "同仁县",
					},
					{
						Name: "尖扎县",
					},
					{
						Name: "泽库县",
					},
					{
						Name: "河南蒙古族自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "海南藏族自治州",
				Children: []*Distinct{
					{
						Name: "共和县",
					},
					{
						Name: "同德县",
					},
					{
						Name: "贵德县",
					},
					{
						Name: "兴海县",
					},
					{
						Name: "贵南县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "果洛藏族自治州",
				Children: []*Distinct{
					{
						Name: "玛沁县",
					},
					{
						Name: "班玛县",
					},
					{
						Name: "甘德县",
					},
					{
						Name: "达日县",
					},
					{
						Name: "久治县",
					},
					{
						Name: "玛多县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "玉树藏族自治州",
				Children: []*Distinct{
					{
						Name: "玉树市",
					},
					{
						Name: "杂多县",
					},
					{
						Name: "称多县",
					},
					{
						Name: "治多县",
					},
					{
						Name: "囊谦县",
					},
					{
						Name: "曲麻莱县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "海西蒙古族藏族自治州",
				Children: []*Distinct{
					{
						Name: "格尔木市",
					},
					{
						Name: "德令哈市",
					},
					{
						Name: "乌兰县",
					},
					{
						Name: "都兰县",
					},
					{
						Name: "天峻县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "宁夏回族自治区",
		Children: []*City{
			{
				Name: "银川市",
				Children: []*Distinct{
					{
						Name: "兴庆区",
					},
					{
						Name: "西夏区",
					},
					{
						Name: "金凤区",
					},
					{
						Name: "永宁县",
					},
					{
						Name: "贺兰县",
					},
					{
						Name: "灵武市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "石嘴山市",
				Children: []*Distinct{
					{
						Name: "大武口区",
					},
					{
						Name: "惠农区",
					},
					{
						Name: "平罗县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "吴忠市",
				Children: []*Distinct{
					{
						Name: "利通区",
					},
					{
						Name: "红寺堡区",
					},
					{
						Name: "盐池县",
					},
					{
						Name: "同心县",
					},
					{
						Name: "青铜峡市",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "固原市",
				Children: []*Distinct{
					{
						Name: "原州区",
					},
					{
						Name: "西吉县",
					},
					{
						Name: "隆德县",
					},
					{
						Name: "泾源县",
					},
					{
						Name: "彭阳县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "中卫市",
				Children: []*Distinct{
					{
						Name: "沙坡头区",
					},
					{
						Name: "中宁县",
					},
					{
						Name: "海原县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "新疆维吾尔自治区",
		Children: []*City{
			{
				Name: "乌鲁木齐市",
				Children: []*Distinct{
					{
						Name: "天山区",
					},
					{
						Name: "沙依巴克区",
					},
					{
						Name: "新市区",
					},
					{
						Name: "水磨沟区",
					},
					{
						Name: "头屯河区",
					},
					{
						Name: "达坂城区",
					},
					{
						Name: "米东区",
					},
					{
						Name: "乌鲁木齐县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "克拉玛依市",
				Children: []*Distinct{
					{
						Name: "独山子区",
					},
					{
						Name: "克拉玛依区",
					},
					{
						Name: "白碱滩区",
					},
					{
						Name: "乌尔禾区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "吐鲁番地区",
				Children: []*Distinct{
					{
						Name: "吐鲁番市",
					},
					{
						Name: "鄯善县",
					},
					{
						Name: "托克逊县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "哈密地区",
				Children: []*Distinct{
					{
						Name: "哈密市",
					},
					{
						Name: "巴里坤哈萨克自治县",
					},
					{
						Name: "伊吾县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "昌吉回族自治州",
				Children: []*Distinct{
					{
						Name: "昌吉市",
					},
					{
						Name: "阜康市",
					},
					{
						Name: "呼图壁县",
					},
					{
						Name: "玛纳斯县",
					},
					{
						Name: "奇台县",
					},
					{
						Name: "吉木萨尔县",
					},
					{
						Name: "木垒哈萨克自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "博尔塔拉蒙古自治州",
				Children: []*Distinct{
					{
						Name: "博乐市",
					},
					{
						Name: "阿拉山口市",
					},
					{
						Name: "精河县",
					},
					{
						Name: "温泉县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "巴音郭楞蒙古自治州",
				Children: []*Distinct{
					{
						Name: "库尔勒市",
					},
					{
						Name: "轮台县",
					},
					{
						Name: "尉犁县",
					},
					{
						Name: "若羌县",
					},
					{
						Name: "且末县",
					},
					{
						Name: "焉耆回族自治县",
					},
					{
						Name: "和静县",
					},
					{
						Name: "和硕县",
					},
					{
						Name: "博湖县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "阿克苏地区",
				Children: []*Distinct{
					{
						Name: "阿克苏市",
					},
					{
						Name: "温宿县",
					},
					{
						Name: "库车县",
					},
					{
						Name: "沙雅县",
					},
					{
						Name: "新和县",
					},
					{
						Name: "拜城县",
					},
					{
						Name: "乌什县",
					},
					{
						Name: "阿瓦提县",
					},
					{
						Name: "柯坪县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "克孜勒苏柯尔克孜自治州",
				Children: []*Distinct{
					{
						Name: "阿图什市",
					},
					{
						Name: "阿克陶县",
					},
					{
						Name: "阿合奇县",
					},
					{
						Name: "乌恰县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "喀什地区",
				Children: []*Distinct{
					{
						Name: "喀什市",
					},
					{
						Name: "疏附县",
					},
					{
						Name: "疏勒县",
					},
					{
						Name: "英吉沙县",
					},
					{
						Name: "泽普县",
					},
					{
						Name: "莎车县",
					},
					{
						Name: "叶城县",
					},
					{
						Name: "麦盖提县",
					},
					{
						Name: "岳普湖县",
					},
					{
						Name: "伽师县",
					},
					{
						Name: "巴楚县",
					},
					{
						Name: "塔什库尔干塔吉克自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "和田地区",
				Children: []*Distinct{
					{
						Name: "和田市",
					},
					{
						Name: "和田县",
					},
					{
						Name: "墨玉县",
					},
					{
						Name: "皮山县",
					},
					{
						Name: "洛浦县",
					},
					{
						Name: "策勒县",
					},
					{
						Name: "于田县",
					},
					{
						Name: "民丰县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "伊犁哈萨克自治州",
				Children: []*Distinct{
					{
						Name: "伊宁市",
					},
					{
						Name: "奎屯市",
					},
					{
						Name: "伊宁县",
					},
					{
						Name: "察布查尔锡伯自治县",
					},
					{
						Name: "霍城县",
					},
					{
						Name: "巩留县",
					},
					{
						Name: "新源县",
					},
					{
						Name: "昭苏县",
					},
					{
						Name: "特克斯县",
					},
					{
						Name: "尼勒克县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "塔城地区",
				Children: []*Distinct{
					{
						Name: "塔城市",
					},
					{
						Name: "乌苏市",
					},
					{
						Name: "额敏县",
					},
					{
						Name: "沙湾县",
					},
					{
						Name: "托里县",
					},
					{
						Name: "裕民县",
					},
					{
						Name: "和布克赛尔蒙古自治县",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "阿勒泰地区",
				Children: []*Distinct{
					{
						Name: "阿勒泰市",
					},
					{
						Name: "布尔津县",
					},
					{
						Name: "富蕴县",
					},
					{
						Name: "福海县",
					},
					{
						Name: "哈巴河县",
					},
					{
						Name: "青河县",
					},
					{
						Name: "吉木乃县",
					},
					{
						Name: "其它区",
					},
				},
			},
		},
	},
	{
		Name: "台湾",
		Children: []*City{
			{
				Name: "台北市",
				Children: []*Distinct{
					{
						Name: "中正区",
					},
					{
						Name: "大同区",
					},
					{
						Name: "中山区",
					},
					{
						Name: "松山区",
					},
					{
						Name: "大安区",
					},
					{
						Name: "万华区",
					},
					{
						Name: "信义区",
					},
					{
						Name: "士林区",
					},
					{
						Name: "北投区",
					},
					{
						Name: "内湖区",
					},
					{
						Name: "南港区",
					},
					{
						Name: "文山区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "高雄市",
				Children: []*Distinct{
					{
						Name: "新兴区",
					},
					{
						Name: "前金区",
					},
					{
						Name: "芩雅区",
					},
					{
						Name: "盐埕区",
					},
					{
						Name: "鼓山区",
					},
					{
						Name: "旗津区",
					},
					{
						Name: "前镇区",
					},
					{
						Name: "三民区",
					},
					{
						Name: "左营区",
					},
					{
						Name: "楠梓区",
					},
					{
						Name: "小港区",
					},
					{
						Name: "其它区",
					},
					{
						Name: "苓雅区",
					},
					{
						Name: "仁武区",
					},
					{
						Name: "大社区",
					},
					{
						Name: "冈山区",
					},
					{
						Name: "路竹区",
					},
					{
						Name: "阿莲区",
					},
					{
						Name: "田寮区",
					},
					{
						Name: "燕巢区",
					},
					{
						Name: "桥头区",
					},
					{
						Name: "梓官区",
					},
					{
						Name: "弥陀区",
					},
					{
						Name: "永安区",
					},
					{
						Name: "湖内区",
					},
					{
						Name: "凤山区",
					},
					{
						Name: "大寮区",
					},
					{
						Name: "林园区",
					},
					{
						Name: "鸟松区",
					},
					{
						Name: "大树区",
					},
					{
						Name: "旗山区",
					},
					{
						Name: "美浓区",
					},
					{
						Name: "六龟区",
					},
					{
						Name: "内门区",
					},
					{
						Name: "杉林区",
					},
					{
						Name: "甲仙区",
					},
					{
						Name: "桃源区",
					},
					{
						Name: "那玛夏区",
					},
					{
						Name: "茂林区",
					},
					{
						Name: "茄萣区",
					},
				},
			},
			{
				Name: "台南市",
				Children: []*Distinct{
					{
						Name: "中西区",
					},
					{
						Name: "东区",
					},
					{
						Name: "南区",
					},
					{
						Name: "北区",
					},
					{
						Name: "安平区",
					},
					{
						Name: "安南区",
					},
					{
						Name: "其它区",
					},
					{
						Name: "永康区",
					},
					{
						Name: "归仁区",
					},
					{
						Name: "新化区",
					},
					{
						Name: "左镇区",
					},
					{
						Name: "玉井区",
					},
					{
						Name: "楠西区",
					},
					{
						Name: "南化区",
					},
					{
						Name: "仁德区",
					},
					{
						Name: "关庙区",
					},
					{
						Name: "龙崎区",
					},
					{
						Name: "官田区",
					},
					{
						Name: "麻豆区",
					},
					{
						Name: "佳里区",
					},
					{
						Name: "西港区",
					},
					{
						Name: "七股区",
					},
					{
						Name: "将军区",
					},
					{
						Name: "学甲区",
					},
					{
						Name: "北门区",
					},
					{
						Name: "新营区",
					},
					{
						Name: "后壁区",
					},
					{
						Name: "白河区",
					},
					{
						Name: "东山区",
					},
					{
						Name: "六甲区",
					},
					{
						Name: "下营区",
					},
					{
						Name: "柳营区",
					},
					{
						Name: "盐水区",
					},
					{
						Name: "善化区",
					},
					{
						Name: "大内区",
					},
					{
						Name: "山上区",
					},
					{
						Name: "新市区",
					},
					{
						Name: "安定区",
					},
				},
			},
			{
				Name: "台中市",
				Children: []*Distinct{
					{
						Name: "中区",
					},
					{
						Name: "东区",
					},
					{
						Name: "南区",
					},
					{
						Name: "西区",
					},
					{
						Name: "北区",
					},
					{
						Name: "北屯区",
					},
					{
						Name: "西屯区",
					},
					{
						Name: "南屯区",
					},
					{
						Name: "其它区",
					},
					{
						Name: "太平区",
					},
					{
						Name: "大里区",
					},
					{
						Name: "雾峰区",
					},
					{
						Name: "乌日区",
					},
					{
						Name: "丰原区",
					},
					{
						Name: "后里区",
					},
					{
						Name: "石冈区",
					},
					{
						Name: "东势区",
					},
					{
						Name: "和平区",
					},
					{
						Name: "新社区",
					},
					{
						Name: "潭子区",
					},
					{
						Name: "大雅区",
					},
					{
						Name: "神冈区",
					},
					{
						Name: "大肚区",
					},
					{
						Name: "沙鹿区",
					},
					{
						Name: "龙井区",
					},
					{
						Name: "梧栖区",
					},
					{
						Name: "清水区",
					},
					{
						Name: "大甲区",
					},
					{
						Name: "外埔区",
					},
					{
						Name: "大安区",
					},
				},
			},
			{
				Name: "金门县",
				Children: []*Distinct{
					{
						Name: "金沙镇",
					},
					{
						Name: "金湖镇",
					},
					{
						Name: "金宁乡",
					},
					{
						Name: "金城镇",
					},
					{
						Name: "烈屿乡",
					},
					{
						Name: "乌坵乡",
					},
				},
			},
			{
				Name: "南投县",
				Children: []*Distinct{
					{
						Name: "南投市",
					},
					{
						Name: "中寮乡",
					},
					{
						Name: "草屯镇",
					},
					{
						Name: "国姓乡",
					},
					{
						Name: "埔里镇",
					},
					{
						Name: "仁爱乡",
					},
					{
						Name: "名间乡",
					},
					{
						Name: "集集镇",
					},
					{
						Name: "水里乡",
					},
					{
						Name: "鱼池乡",
					},
					{
						Name: "信义乡",
					},
					{
						Name: "竹山镇",
					},
					{
						Name: "鹿谷乡",
					},
				},
			},
			{
				Name: "基隆市",
				Children: []*Distinct{
					{
						Name: "仁爱区",
					},
					{
						Name: "信义区",
					},
					{
						Name: "中正区",
					},
					{
						Name: "中山区",
					},
					{
						Name: "安乐区",
					},
					{
						Name: "暖暖区",
					},
					{
						Name: "七堵区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "新竹市",
				Children: []*Distinct{
					{
						Name: "东区",
					},
					{
						Name: "北区",
					},
					{
						Name: "香山区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "嘉义市",
				Children: []*Distinct{
					{
						Name: "东区",
					},
					{
						Name: "西区",
					},
					{
						Name: "其它区",
					},
				},
			},
			{
				Name: "新北市",
				Children: []*Distinct{
					{
						Name: "万里区",
					},
					{
						Name: "金山区",
					},
					{
						Name: "板桥区",
					},
					{
						Name: "汐止区",
					},
					{
						Name: "深坑区",
					},
					{
						Name: "石碇区",
					},
					{
						Name: "瑞芳区",
					},
					{
						Name: "平溪区",
					},
					{
						Name: "双溪区",
					},
					{
						Name: "贡寮区",
					},
					{
						Name: "新店区",
					},
					{
						Name: "坪林区",
					},
					{
						Name: "乌来区",
					},
					{
						Name: "永和区",
					},
					{
						Name: "中和区",
					},
					{
						Name: "土城区",
					},
					{
						Name: "三峡区",
					},
					{
						Name: "树林区",
					},
					{
						Name: "莺歌区",
					},
					{
						Name: "三重区",
					},
					{
						Name: "新庄区",
					},
					{
						Name: "泰山区",
					},
					{
						Name: "林口区",
					},
					{
						Name: "芦洲区",
					},
					{
						Name: "五股区",
					},
					{
						Name: "八里区",
					},
					{
						Name: "淡水区",
					},
					{
						Name: "三芝区",
					},
					{
						Name: "石门区",
					},
				},
			},
			{
				Name: "宜兰县",
				Children: []*Distinct{
					{
						Name: "宜兰市",
					},
					{
						Name: "头城镇",
					},
					{
						Name: "礁溪乡",
					},
					{
						Name: "壮围乡",
					},
					{
						Name: "员山乡",
					},
					{
						Name: "罗东镇",
					},
					{
						Name: "三星乡",
					},
					{
						Name: "大同乡",
					},
					{
						Name: "五结乡",
					},
					{
						Name: "冬山乡",
					},
					{
						Name: "苏澳镇",
					},
					{
						Name: "南澳乡",
					},
					{
						Name: "钓鱼台",
					},
				},
			},
			{
				Name: "新竹县",
				Children: []*Distinct{
					{
						Name: "竹北市",
					},
					{
						Name: "湖口乡",
					},
					{
						Name: "新丰乡",
					},
					{
						Name: "新埔镇",
					},
					{
						Name: "关西镇",
					},
					{
						Name: "芎林乡",
					},
					{
						Name: "宝山乡",
					},
					{
						Name: "竹东镇",
					},
					{
						Name: "五峰乡",
					},
					{
						Name: "横山乡",
					},
					{
						Name: "尖石乡",
					},
					{
						Name: "北埔乡",
					},
					{
						Name: "峨眉乡",
					},
				},
			},
			{
				Name: "桃园县",
				Children: []*Distinct{
					{
						Name: "中坜市",
					},
					{
						Name: "平镇市",
					},
					{
						Name: "龙潭乡",
					},
					{
						Name: "杨梅市",
					},
					{
						Name: "新屋乡",
					},
					{
						Name: "观音乡",
					},
					{
						Name: "桃园市",
					},
					{
						Name: "龟山乡",
					},
					{
						Name: "八德市",
					},
					{
						Name: "大溪镇",
					},
					{
						Name: "复兴乡",
					},
					{
						Name: "大园乡",
					},
					{
						Name: "芦竹乡",
					},
				},
			},
			{
				Name: "苗栗县",
				Children: []*Distinct{
					{
						Name: "竹南镇",
					},
					{
						Name: "头份镇",
					},
					{
						Name: "三湾乡",
					},
					{
						Name: "南庄乡",
					},
					{
						Name: "狮潭乡",
					},
					{
						Name: "后龙镇",
					},
					{
						Name: "通霄镇",
					},
					{
						Name: "苑里镇",
					},
					{
						Name: "苗栗市",
					},
					{
						Name: "造桥乡",
					},
					{
						Name: "头屋乡",
					},
					{
						Name: "公馆乡",
					},
					{
						Name: "大湖乡",
					},
					{
						Name: "泰安乡",
					},
					{
						Name: "铜锣乡",
					},
					{
						Name: "三义乡",
					},
					{
						Name: "西湖乡",
					},
					{
						Name: "卓兰镇",
					},
				},
			},
			{
				Name: "彰化县",
				Children: []*Distinct{
					{
						Name: "彰化市",
					},
					{
						Name: "芬园乡",
					},
					{
						Name: "花坛乡",
					},
					{
						Name: "秀水乡",
					},
					{
						Name: "鹿港镇",
					},
					{
						Name: "福兴乡",
					},
					{
						Name: "线西乡",
					},
					{
						Name: "和美镇",
					},
					{
						Name: "伸港乡",
					},
					{
						Name: "员林镇",
					},
					{
						Name: "社头乡",
					},
					{
						Name: "永靖乡",
					},
					{
						Name: "埔心乡",
					},
					{
						Name: "溪湖镇",
					},
					{
						Name: "大村乡",
					},
					{
						Name: "埔盐乡",
					},
					{
						Name: "田中镇",
					},
					{
						Name: "北斗镇",
					},
					{
						Name: "田尾乡",
					},
					{
						Name: "埤头乡",
					},
					{
						Name: "溪州乡",
					},
					{
						Name: "竹塘乡",
					},
					{
						Name: "二林镇",
					},
					{
						Name: "大城乡",
					},
					{
						Name: "芳苑乡",
					},
					{
						Name: "二水乡",
					},
				},
			},
			{
				Name: "嘉义县",
				Children: []*Distinct{
					{
						Name: "番路乡",
					},
					{
						Name: "梅山乡",
					},
					{
						Name: "竹崎乡",
					},
					{
						Name: "阿里山乡",
					},
					{
						Name: "中埔乡",
					},
					{
						Name: "大埔乡",
					},
					{
						Name: "水上乡",
					},
					{
						Name: "鹿草乡",
					},
					{
						Name: "太保市",
					},
					{
						Name: "朴子市",
					},
					{
						Name: "东石乡",
					},
					{
						Name: "六脚乡",
					},
					{
						Name: "新港乡",
					},
					{
						Name: "民雄乡",
					},
					{
						Name: "大林镇",
					},
					{
						Name: "溪口乡",
					},
					{
						Name: "义竹乡",
					},
					{
						Name: "布袋镇",
					},
				},
			},
			{
				Name: "云林县",
				Children: []*Distinct{
					{
						Name: "斗南镇",
					},
					{
						Name: "大埤乡",
					},
					{
						Name: "虎尾镇",
					},
					{
						Name: "土库镇",
					},
					{
						Name: "褒忠乡",
					},
					{
						Name: "东势乡",
					},
					{
						Name: "台西乡",
					},
					{
						Name: "仑背乡",
					},
					{
						Name: "麦寮乡",
					},
					{
						Name: "斗六市",
					},
					{
						Name: "林内乡",
					},
					{
						Name: "古坑乡",
					},
					{
						Name: "莿桐乡",
					},
					{
						Name: "西螺镇",
					},
					{
						Name: "二仑乡",
					},
					{
						Name: "北港镇",
					},
					{
						Name: "水林乡",
					},
					{
						Name: "口湖乡",
					},
					{
						Name: "四湖乡",
					},
					{
						Name: "元长乡",
					},
				},
			},
			{
				Name: "屏东县",
				Children: []*Distinct{
					{
						Name: "屏东市",
					},
					{
						Name: "三地门乡",
					},
					{
						Name: "雾台乡",
					},
					{
						Name: "玛家乡",
					},
					{
						Name: "九如乡",
					},
					{
						Name: "里港乡",
					},
					{
						Name: "高树乡",
					},
					{
						Name: "盐埔乡",
					},
					{
						Name: "长治乡",
					},
					{
						Name: "麟洛乡",
					},
					{
						Name: "竹田乡",
					},
					{
						Name: "内埔乡",
					},
					{
						Name: "万丹乡",
					},
					{
						Name: "潮州镇",
					},
					{
						Name: "泰武乡",
					},
					{
						Name: "来义乡",
					},
					{
						Name: "万峦乡",
					},
					{
						Name: "崁顶乡",
					},
					{
						Name: "新埤乡",
					},
					{
						Name: "南州乡",
					},
					{
						Name: "林边乡",
					},
					{
						Name: "东港镇",
					},
					{
						Name: "琉球乡",
					},
					{
						Name: "佳冬乡",
					},
					{
						Name: "新园乡",
					},
					{
						Name: "枋寮乡",
					},
					{
						Name: "枋山乡",
					},
					{
						Name: "春日乡",
					},
					{
						Name: "狮子乡",
					},
					{
						Name: "车城乡",
					},
					{
						Name: "牡丹乡",
					},
					{
						Name: "恒春镇",
					},
					{
						Name: "满州乡",
					},
				},
			},
			{
				Name: "台东县",
				Children: []*Distinct{
					{
						Name: "台东市",
					},
					{
						Name: "绿岛乡",
					},
					{
						Name: "兰屿乡",
					},
					{
						Name: "延平乡",
					},
					{
						Name: "卑南乡",
					},
					{
						Name: "鹿野乡",
					},
					{
						Name: "关山镇",
					},
					{
						Name: "海端乡",
					},
					{
						Name: "池上乡",
					},
					{
						Name: "东河乡",
					},
					{
						Name: "成功镇",
					},
					{
						Name: "长滨乡",
					},
					{
						Name: "金峰乡",
					},
					{
						Name: "大武乡",
					},
					{
						Name: "达仁乡",
					},
					{
						Name: "太麻里乡",
					},
				},
			},
			{
				Name: "花莲县",
				Children: []*Distinct{
					{
						Name: "花莲市",
					},
					{
						Name: "新城乡",
					},
					{
						Name: "太鲁阁",
					},
					{
						Name: "秀林乡",
					},
					{
						Name: "吉安乡",
					},
					{
						Name: "寿丰乡",
					},
					{
						Name: "凤林镇",
					},
					{
						Name: "光复乡",
					},
					{
						Name: "丰滨乡",
					},
					{
						Name: "瑞穗乡",
					},
					{
						Name: "万荣乡",
					},
					{
						Name: "玉里镇",
					},
					{
						Name: "卓溪乡",
					},
					{
						Name: "富里乡",
					},
				},
			},
			{
				Name: "澎湖县",
				Children: []*Distinct{
					{
						Name: "马公市",
					},
					{
						Name: "西屿乡",
					},
					{
						Name: "望安乡",
					},
					{
						Name: "七美乡",
					},
					{
						Name: "白沙乡",
					},
					{
						Name: "湖西乡",
					},
				},
			},
			{
				Name: "连江县",
				Children: []*Distinct{
					{
						Name: "南竿乡",
					},
					{
						Name: "北竿乡",
					},
					{
						Name: "莒光乡",
					},
					{
						Name: "东引乡",
					},
				},
			},
		},
	},
	{
		Name: "香港特别行政区",
		Children: []*City{
			{
				Name: "香港岛",
				Children: []*Distinct{
					{
						Name: "中西区",
					},
					{
						Name: "湾仔",
					},
					{
						Name: "东区",
					},
					{
						Name: "南区",
					},
				},
			},
			{
				Name: "九龙",
				Children: []*Distinct{
					{
						Name: "九龙城区",
					},
					{
						Name: "油尖旺区",
					},
					{
						Name: "深水埗区",
					},
					{
						Name: "黄大仙区",
					},
					{
						Name: "观塘区",
					},
				},
			},
			{
				Name: "新界",
				Children: []*Distinct{
					{
						Name: "北区",
					},
					{
						Name: "大埔区",
					},
					{
						Name: "沙田区",
					},
					{
						Name: "西贡区",
					},
					{
						Name: "元朗区",
					},
					{
						Name: "屯门区",
					},
					{
						Name: "荃湾区",
					},
					{
						Name: "葵青区",
					},
					{
						Name: "离岛区",
					},
				},
			},
		},
	},
	{
		Name: "澳门特别行政区",
		Children: []*City{
			{
				Name: "澳门半岛",
			},
			{
				Name: "离岛",
			},
		},
	},
	{
		Name: "海外",
		Children: []*City{
			{
				Name: "海外",
			},
		},
	},
}
