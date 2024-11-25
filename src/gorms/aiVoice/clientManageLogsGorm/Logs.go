package clientManageLogsGorm

type Table struct {
	Id                   int    `gorm:"column:Id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"Id"`
	CreatedAt            string `gorm:"column:CreatedAt;type:datetime" json:"CreatedAt"`
	UpdatedAt            string `gorm:"column:UpdatedAt;type:datetime" json:"UpdatedAt"`
	CreatedName          string `gorm:"column:CreatedName;type:varchar(50)" json:"CreatedName"`
	RecordId             int    `gorm:"column:RecordId;type:int(10)" json:"RecordId"`
	Callere164           string `gorm:"column:Callere164;type:varchar(200)" json:"Callere164"`
	Calleraccesse164     string `gorm:"column:Calleraccesse164;type:varchar(200)" json:"Calleraccesse164"`
	Calleee164           string `gorm:"column:Calleee164;type:varchar(200)" json:"Calleee164"`
	Calleeaccesse164     string `gorm:"column:Calleeaccesse164;type:varchar(200)" json:"Calleeaccesse164"`
	Callerip             string `gorm:"column:Callerip;type:varchar(200)" json:"Callerip"`
	Callerrtpip          string `gorm:"column:Callerrtpip;type:varchar(200)" json:"Callerrtpip"`
	Callercodec          string `gorm:"column:Callercodec;type:varchar(200)" json:"Callercodec"`
	Callergatewayid      string `gorm:"column:Callergatewayid;type:varchar(200)" json:"Callergatewayid"`
	Callerproductid      string `gorm:"column:Callerproductid;type:varchar(200)" json:"Callerproductid"`
	Callertogatewaye164  string `gorm:"column:Callertogatewaye164;type:varchar(200)" json:"Callertogatewaye164"`
	Callertype           int    `gorm:"column:Callertype;type:int(10)" json:"Callertype"`
	Calleeip             string `gorm:"column:Calleeip;type:varchar(200)" json:"Calleeip"`
	Calleertpip          string `gorm:"column:Calleertpip;type:varchar(200)" json:"Calleertpip"`
	Calleecodec          string `gorm:"column:Calleecodec;type:varchar(200)" json:"Calleecodec"`
	Calleegatewayid      string `gorm:"column:Calleegatewayid;type:varchar(200)" json:"Calleegatewayid"`
	Calleeproductid      string `gorm:"column:Calleeproductid;type:varchar(200)" json:"Calleeproductid"`
	Calleetogatewaye164  string `gorm:"column:Calleetogatewaye164;type:varchar(200)" json:"Calleetogatewaye164"`
	Calleetype           int    `gorm:"column:Calleetype;type:int(10)" json:"Calleetype"`
	Billingmode          int    `gorm:"column:Billingmode;type:int(10)" json:"Billingmode"`
	Calllevel            int    `gorm:"column:Calllevel;type:int(10)" json:"Calllevel"`
	Agentfeetime         int    `gorm:"column:Agentfeetime;type:int(10)" json:"Agentfeetime"`
	Starttime            int64  `gorm:"column:Starttime;type:bigint(20)" json:"Starttime"`
	Stoptime             string `gorm:"column:Stoptime;type:varchar(20)" json:"Stoptime"`
	Callerpdd            int    `gorm:"column:Callerpdd;type:int(10)" json:"Callerpdd"`
	Calleepdd            int    `gorm:"column:Calleepdd;type:int(10)" json:"Calleepdd"`
	Holdtime             int    `gorm:"column:Holdtime;type:int(10)" json:"Holdtime"`
	Callerareacode       string `gorm:"column:Callerareacode;type:varchar(200)" json:"Callerareacode"`
	Feetime              int    `gorm:"column:Feetime;type:int(10)" json:"Feetime"`
	Fee                  string `gorm:"column:Fee;type:varchar(50)" json:"Fee"`
	Tax                  string `gorm:"column:Tax;type:varchar(50)" json:"Tax"`
	Suitefee             string `gorm:"column:Suitefee;type:varchar(50)" json:"Suitefee"`
	Suitefeetime         int    `gorm:"column:Suitefeetime;type:int(10)" json:"Suitefeetime"`
	Incomefee            string `gorm:"column:Incomefee;type:varchar(200)" json:"Incomefee"`
	Incometax            string `gorm:"column:Incometax;type:varchar(200)" json:"Incometax"`
	Customeraccount      string `gorm:"column:Customeraccount;type:varchar(200)" json:"Customeraccount"`
	Customername         string `gorm:"column:Customername;type:varchar(200)" json:"Customername"`
	Calleeareacode       string `gorm:"column:Calleeareacode;type:varchar(200)" json:"Calleeareacode"`
	Agentfee             string `gorm:"column:Agentfee;type:varchar(200)" json:"Agentfee"`
	Agenttax             string `gorm:"column:Agenttax;type:varchar(200)" json:"Agenttax"`
	Agentsuitefee        string `gorm:"column:Agentsuitefee;type:varchar(200)" json:"Agentsuitefee"`
	Agentsuitefeetime    int    `gorm:"column:Agentsuitefeetime;type:int(10)" json:"Agentsuitefeetime"`
	Agentaccount         string `gorm:"column:Agentaccount;type:varchar(200)" json:"Agentaccount"`
	Agentname            string `gorm:"column:Agentname;type:varchar(200)" json:"Agentname"`
	Flowno               string `gorm:"column:Flowno;type:varchar(200)" json:"Flowno"`
	Softswitchname       string `gorm:"column:Softswitchname;type:varchar(200)" json:"Softswitchname"`
	Softswitchcallid     string `gorm:"column:Softswitchcallid;type:varchar(200)" json:"Softswitchcallid"`
	Callercallid         string `gorm:"column:Callercallid;type:varchar(200)" json:"Callercallid"`
	Calleroriginalcallid string `gorm:"column:Calleroriginalcallid;type:varchar(200)" json:"Calleroriginalcallid"`
	Calleecallid         string `gorm:"column:Calleecallid;type:varchar(200)" json:"Calleecallid"`
	Calleroriginalinfo   string `gorm:"column:Calleroriginalinfo;type:varchar(200)" json:"Calleroriginalinfo"`
	Rtpforward           int    `gorm:"column:Rtpforward;type:int(10)" json:"Rtpforward"`
	Enddirection         int    `gorm:"column:Enddirection;type:int(10)" json:"Enddirection"`
	Endreason            int    `gorm:"column:Endreason;type:int(10)" json:"Endreason"`
	Billingtype          int    `gorm:"column:Billingtype;type:int(10)" json:"Billingtype"`
	Cdrlevel             int    `gorm:"column:Cdrlevel;type:int(10)" json:"Cdrlevel"`
	AgentcdrId           int    `gorm:"column:AgentcdrId;type:int(10)" json:"AgentcdrId"`
	Sipreasonheader      string `gorm:"column:Sipreasonheader;type:varchar(200)" json:"Sipreasonheader"`
	Recordstarttime      string `gorm:"column:Recordstarttime;type:varchar(200)" json:"Recordstarttime"`
	Transactionid        string `gorm:"column:Transactionid;type:varchar(200)" json:"Transactionid"`
	Flownofirst          string `gorm:"column:Flownofirst;type:varchar(200)" json:"Flownofirst"`
	Additional           string `gorm:"column:Additional;type:text" json:"Additional"`
	DynamicValue         string `gorm:"column:DynamicValue;type:text" json:"DynamicValue"`
	TrieResult           string `gorm:"column:TrieResult;type:json" json:"TrieResult"`
	LeftResult           string `gorm:"column:LeftResult;type:longtext" json:"LeftResult"`
	RightResult          string `gorm:"column:RightResult;type:longtext" json:"RightResult"`
	RecordUrl            string `gorm:"column:RecordUrl;type:text" json:"RecordUrl"`
	IndustryTrie         string `gorm:"column:IndustryTrie;type:text" json:"IndustryTrie"`
	ClientTrie           string `gorm:"column:ClientTrie;type:text" json:"ClientTrie"`
	ClientManageRuleId   int    `gorm:"column:ClientManageRuleId;type:int(10)" json:"ClientManageRuleId"`
	IP                   string `gorm:"column:IP;type:varchar(20)" json:"IP"`
	OssUrl               string `gorm:"column:OssUrl;type:text" json:"OssUrl"`
	SopResult            string `gorm:"column:SopResult;type:text" json:"SopResult"`
	SopStatus            int    `gorm:"column:SopStatus;type:int(10)" map:"ClientManageLogsSopStatusMap" json:"SopStatus"`
	Status               int    `gorm:"column:Status;type:int(10)" map:"ClientManageLogsStatusMap" json:"Status"`
	ErrMsg               string `gorm:"column:ErrMsg;type:text" json:"ErrMsg"`
	AuditStatus          int    `gorm:"column:AuditStatus;type:int(10)" map:"ClientManageLogsAuditStatusMap" json:"AuditStatus"`
	CostSecond           int    `gorm:"column:CostSecond;type:int(10)" json:"CostSecond"`
}

func (m *Table) TableName() string {
	return "clientManageLogs"
}
