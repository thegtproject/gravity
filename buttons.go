package gravity

// Button ...
type Button int

// String ...
func (b Button) String() string {
	name, ok := buttonNames[b]
	if !ok {
		return "Invalid"
	}
	return name
}

// List of all mouse buttons.
const (
	MouseButton1 = Button(0)
	MouseButton2 = Button(1)
	MouseButton3 = Button(2)
	MouseButton4 = Button(3)
	MouseLeft    = Button(MouseButton1)
	MouseRight   = Button(MouseButton2)
)

// List of all keyboard buttons.
const (
	KeySpace        = Button(32)
	KeyApostrophe   = Button(39)
	KeyComma        = Button(44)
	KeyMinus        = Button(45)
	KeyPeriod       = Button(46)
	KeySlash        = Button(47)
	Key0            = Button(48)
	Key1            = Button(49)
	Key2            = Button(50)
	Key3            = Button(51)
	Key4            = Button(52)
	Key5            = Button(53)
	Key6            = Button(54)
	Key7            = Button(55)
	Key8            = Button(56)
	Key9            = Button(57)
	KeySemicolon    = Button(59)
	KeyEqual        = Button(61)
	KeyA            = Button(65)
	KeyB            = Button(66)
	KeyC            = Button(67)
	KeyD            = Button(68)
	KeyE            = Button(69)
	KeyF            = Button(70)
	KeyG            = Button(71)
	KeyH            = Button(72)
	KeyI            = Button(73)
	KeyJ            = Button(74)
	KeyK            = Button(75)
	KeyL            = Button(76)
	KeyM            = Button(77)
	KeyN            = Button(78)
	KeyO            = Button(79)
	KeyP            = Button(80)
	KeyQ            = Button(81)
	KeyR            = Button(82)
	KeyS            = Button(83)
	KeyT            = Button(84)
	KeyU            = Button(85)
	KeyV            = Button(86)
	KeyW            = Button(87)
	KeyX            = Button(88)
	KeyY            = Button(89)
	KeyZ            = Button(90)
	KeyLeftBracket  = Button(91)
	KeyBackSlash    = Button(92)
	KeyRightBracket = Button(93)
	KeyGrave        = Button(96)
	KeyEscape       = Button(256)
	KeyEnter        = Button(257)
	KeyTab          = Button(258)
	KeyBackspace    = Button(259)
	KeyInsert       = Button(260)
	KeyDelete       = Button(261)
	KeyRight        = Button(262)
	KeyLeft         = Button(263)
	KeyDown         = Button(264)
	KeyUp           = Button(265)
	KeyPageUp       = Button(266)
	KeyPageDown     = Button(267)
	KeyHome         = Button(268)
	KeyEnd          = Button(269)
	KeyCaps         = Button(280)
	KeyScroll       = Button(281)
	KeyNumlock      = Button(282)
	KeyPrintScreen  = Button(283)
	KeyPause        = Button(284)
	KeyF1           = Button(290)
	KeyF2           = Button(291)
	KeyF3           = Button(292)
	KeyF4           = Button(293)
	KeyF5           = Button(294)
	KeyF6           = Button(295)
	KeyF7           = Button(296)
	KeyF8           = Button(297)
	KeyF9           = Button(298)
	KeyF10          = Button(299)
	KeyF11          = Button(300)
	KeyF12          = Button(301)
	KeyF13          = Button(302)
	KeyF14          = Button(303)
	KeyF15          = Button(304)
	KeyF16          = Button(305)
	KeyF17          = Button(306)
	KeyF18          = Button(307)
	KeyF19          = Button(308)
	KeyF20          = Button(309)
	KeyF21          = Button(310)
	KeyF22          = Button(311)
	KeyF23          = Button(312)
	KeyF24          = Button(313)
	KeyF25          = Button(314)
	KeyKP0          = Button(320)
	KeyKP1          = Button(321)
	KeyKP2          = Button(322)
	KeyKP3          = Button(323)
	KeyKP4          = Button(324)
	KeyKP5          = Button(325)
	KeyKP6          = Button(326)
	KeyKP7          = Button(327)
	KeyKP8          = Button(328)
	KeyKP9          = Button(329)
	KeyDecimal      = Button(330)
	KeyDivide       = Button(331)
	KeyMultiply     = Button(332)
	KeySubtract     = Button(333)
	KeyAdd          = Button(334)
	KeyKPEnter      = Button(335)
	KeyKPEqual      = Button(336)
	KeyLeftShift    = Button(340)
	KeyLeftControl  = Button(341)
	KeyLeftAlt      = Button(342)
	KeyLeftSuper    = Button(343)
	KeyRightShift   = Button(344)
	KeyRightControl = Button(345)
	KeyRightAlt     = Button(346)
	KeyRightSuper   = Button(347)
	KeyMenu         = Button(348)
)

var buttonNames = map[Button]string{
	Button(32):  "KeySpace",
	Button(39):  "KeyApostrophe",
	Button(44):  "KeyComma",
	Button(45):  "KeyMinus",
	Button(46):  "KeyPeriod",
	Button(47):  "KeySlash",
	Button(48):  "Key0",
	Button(49):  "Key1",
	Button(50):  "Key2",
	Button(51):  "Key3",
	Button(52):  "Key4",
	Button(53):  "Key5",
	Button(54):  "Key6",
	Button(55):  "Key7",
	Button(56):  "Key8",
	Button(57):  "Key9",
	Button(59):  "KeySemicolon",
	Button(61):  "KeyEqual",
	Button(65):  "KeyA",
	Button(66):  "KeyB",
	Button(67):  "KeyC",
	Button(68):  "KeyD",
	Button(69):  "KeyE",
	Button(70):  "KeyF",
	Button(71):  "KeyG",
	Button(72):  "KeyH",
	Button(73):  "KeyI",
	Button(74):  "KeyJ",
	Button(75):  "KeyK",
	Button(76):  "KeyL",
	Button(77):  "KeyM",
	Button(78):  "KeyN",
	Button(79):  "KeyO",
	Button(80):  "KeyP",
	Button(81):  "KeyQ",
	Button(82):  "KeyR",
	Button(83):  "KeyS",
	Button(84):  "KeyT",
	Button(85):  "KeyU",
	Button(86):  "KeyV",
	Button(87):  "KeyW",
	Button(88):  "KeyX",
	Button(89):  "KeyY",
	Button(90):  "KeyZ",
	Button(91):  "KeyLeftBracket",
	Button(92):  "KeyBackSlash",
	Button(93):  "KeyRightBracket",
	Button(96):  "KeyGrave",
	Button(256): "KeyEscape",
	Button(257): "KeyEnter",
	Button(258): "KeyTab",
	Button(259): "KeyBackspace",
	Button(260): "KeyInsert",
	Button(261): "KeyDelete",
	Button(262): "KeyRight",
	Button(263): "KeyLeft",
	Button(264): "KeyDown",
	Button(265): "KeyUp",
	Button(266): "KeyPageUp",
	Button(267): "KeyPageDown",
	Button(268): "KeyHome",
	Button(269): "KeyEnd",
	Button(280): "KeyCaps",
	Button(281): "KeyScroll",
	Button(282): "KeyNumlock",
	Button(283): "KeyPrintScreen",
	Button(284): "KeyPause",
	Button(290): "KeyF1",
	Button(291): "KeyF2",
	Button(292): "KeyF3",
	Button(293): "KeyF4",
	Button(294): "KeyF5",
	Button(295): "KeyF6",
	Button(296): "KeyF7",
	Button(297): "KeyF8",
	Button(298): "KeyF9",
	Button(299): "KeyF10",
	Button(300): "KeyF11",
	Button(301): "KeyF12",
	Button(302): "KeyF13",
	Button(303): "KeyF14",
	Button(304): "KeyF15",
	Button(305): "KeyF16",
	Button(306): "KeyF17",
	Button(307): "KeyF18",
	Button(308): "KeyF19",
	Button(309): "KeyF20",
	Button(310): "KeyF21",
	Button(311): "KeyF22",
	Button(312): "KeyF23",
	Button(313): "KeyF24",
	Button(314): "KeyF25",
	Button(320): "KeyKP0",
	Button(321): "KeyKP1",
	Button(322): "KeyKP2",
	Button(323): "KeyKP3",
	Button(324): "KeyKP4",
	Button(325): "KeyKP5",
	Button(326): "KeyKP6",
	Button(327): "KeyKP7",
	Button(328): "KeyKP8",
	Button(329): "KeyKP9",
	Button(330): "KeyDecimal",
	Button(331): "KeyDivide",
	Button(332): "KeyMultiply",
	Button(333): "KeySubtract",
	Button(334): "KeyAdd",
	Button(335): "KeyKPEnter",
	Button(336): "KeyKPEqual",
	Button(340): "KeyLeftShift",
	Button(341): "KeyLeftControl",
	Button(342): "KeyLeftAlt",
	Button(343): "KeyLeftSuper",
	Button(344): "KeyRightShift",
	Button(345): "KeyRightControl",
	Button(346): "KeyRightAlt",
	Button(347): "KeyRightSuper",
	Button(348): "KeyMenu",
}
