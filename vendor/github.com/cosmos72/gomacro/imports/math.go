// this file was generated by gomacro command: import _b "math"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"math"
)

// reflection: allow interpreted code to import "math"
func init() {
	Packages["math"] = Package{
	Binds: map[string]Value{
		"Abs":	ValueOf(math.Abs),
		"Acos":	ValueOf(math.Acos),
		"Acosh":	ValueOf(math.Acosh),
		"Asin":	ValueOf(math.Asin),
		"Asinh":	ValueOf(math.Asinh),
		"Atan":	ValueOf(math.Atan),
		"Atan2":	ValueOf(math.Atan2),
		"Atanh":	ValueOf(math.Atanh),
		"Cbrt":	ValueOf(math.Cbrt),
		"Ceil":	ValueOf(math.Ceil),
		"Copysign":	ValueOf(math.Copysign),
		"Cos":	ValueOf(math.Cos),
		"Cosh":	ValueOf(math.Cosh),
		"Dim":	ValueOf(math.Dim),
		"E":	ValueOf(math.E),
		"Erf":	ValueOf(math.Erf),
		"Erfc":	ValueOf(math.Erfc),
		"Exp":	ValueOf(math.Exp),
		"Exp2":	ValueOf(math.Exp2),
		"Expm1":	ValueOf(math.Expm1),
		"Float32bits":	ValueOf(math.Float32bits),
		"Float32frombits":	ValueOf(math.Float32frombits),
		"Float64bits":	ValueOf(math.Float64bits),
		"Float64frombits":	ValueOf(math.Float64frombits),
		"Floor":	ValueOf(math.Floor),
		"Frexp":	ValueOf(math.Frexp),
		"Gamma":	ValueOf(math.Gamma),
		"Hypot":	ValueOf(math.Hypot),
		"Ilogb":	ValueOf(math.Ilogb),
		"Inf":	ValueOf(math.Inf),
		"IsInf":	ValueOf(math.IsInf),
		"IsNaN":	ValueOf(math.IsNaN),
		"J0":	ValueOf(math.J0),
		"J1":	ValueOf(math.J1),
		"Jn":	ValueOf(math.Jn),
		"Ldexp":	ValueOf(math.Ldexp),
		"Lgamma":	ValueOf(math.Lgamma),
		"Ln10":	ValueOf(math.Ln10),
		"Ln2":	ValueOf(math.Ln2),
		"Log":	ValueOf(math.Log),
		"Log10":	ValueOf(math.Log10),
		"Log10E":	ValueOf(math.Log10E),
		"Log1p":	ValueOf(math.Log1p),
		"Log2":	ValueOf(math.Log2),
		"Log2E":	ValueOf(math.Log2E),
		"Logb":	ValueOf(math.Logb),
		"Max":	ValueOf(math.Max),
		"MaxFloat32":	ValueOf(float32(math.MaxFloat32)),
		"MaxFloat64":	ValueOf(float64(math.MaxFloat64)),
		"MaxInt16":	ValueOf(math.MaxInt16),
		"MaxInt32":	ValueOf(math.MaxInt32),
		"MaxInt64":	ValueOf(int64(math.MaxInt64)),
		"MaxInt8":	ValueOf(math.MaxInt8),
		"MaxUint16":	ValueOf(math.MaxUint16),
		"MaxUint32":	ValueOf(uint32(math.MaxUint32)),
		"MaxUint64":	ValueOf(uint64(math.MaxUint64)),
		"MaxUint8":	ValueOf(math.MaxUint8),
		"Min":	ValueOf(math.Min),
		"MinInt16":	ValueOf(math.MinInt16),
		"MinInt32":	ValueOf(math.MinInt32),
		"MinInt64":	ValueOf(int64(math.MinInt64)),
		"MinInt8":	ValueOf(math.MinInt8),
		"Mod":	ValueOf(math.Mod),
		"Modf":	ValueOf(math.Modf),
		"NaN":	ValueOf(math.NaN),
		"Nextafter":	ValueOf(math.Nextafter),
		"Nextafter32":	ValueOf(math.Nextafter32),
		"Phi":	ValueOf(math.Phi),
		"Pi":	ValueOf(math.Pi),
		"Pow":	ValueOf(math.Pow),
		"Pow10":	ValueOf(math.Pow10),
		"Remainder":	ValueOf(math.Remainder),
		"Signbit":	ValueOf(math.Signbit),
		"Sin":	ValueOf(math.Sin),
		"Sincos":	ValueOf(math.Sincos),
		"Sinh":	ValueOf(math.Sinh),
		"SmallestNonzeroFloat32":	ValueOf(math.SmallestNonzeroFloat32),
		"SmallestNonzeroFloat64":	ValueOf(math.SmallestNonzeroFloat64),
		"Sqrt":	ValueOf(math.Sqrt),
		"Sqrt2":	ValueOf(math.Sqrt2),
		"SqrtE":	ValueOf(math.SqrtE),
		"SqrtPhi":	ValueOf(math.SqrtPhi),
		"SqrtPi":	ValueOf(math.SqrtPi),
		"Tan":	ValueOf(math.Tan),
		"Tanh":	ValueOf(math.Tanh),
		"Trunc":	ValueOf(math.Trunc),
		"Y0":	ValueOf(math.Y0),
		"Y1":	ValueOf(math.Y1),
		"Yn":	ValueOf(math.Yn),
	},Untypeds: map[string]string{
		"E":	"float:9111550163858012281440901732746538838772262590143654133938674751170893736363860704678356685906435473285900222617098459660313571825500424586151709661124231/3351951982485649274893506249551461531869841455148098344430890360930441007518386744200468574541725856922507964546621512713438470702986642486608412251521024",
		"Ln10":	"float:7718154667303294525535807123668701784088749544639269844330854713417502399132378792470215254015532173856280403153541607081951593465883977341501885089972609/3351951982485649274893506249551461531869841455148098344430890360930441007518386744200468574541725856922507964546621512713438470702986642486608412251521024",
		"Ln2":	"float:1161698033016123487673055082581635139703829175621884955134501220495379707050587855317621548207870502811545419056917762934492002063302542007440018497053197/1675975991242824637446753124775730765934920727574049172215445180465220503759193372100234287270862928461253982273310756356719235351493321243304206125760512",
		"Log10E":	"float:11645873996785463169335113650913961638081378781380835843607526526887199042223674282440048773139860374459897757791791429659694786306613464090293660335282213/26815615859885194199148049996411692254958731641184786755447122887443528060147093953603748596333806855380063716372972101707507765623893139892867298012168192",
		"Log2E":	"float:9671689004859951471451510700966227821397578809549890556103133486120211484101236374917210827834571490990560020909793651977303669442567160330196651027673915/6703903964971298549787012499102923063739682910296196688861780721860882015036773488400937149083451713845015929093243025426876941405973284973216824503042048",
		"MaxFloat32":	"float:340282346638528859811704183484516925440",
		"MaxFloat64":	"float:179769313486231570814527423731704356798100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001726091719747851502452248840877809013992353692576760003314411346346787114201583914540245118557675136988317857687302483414838584368727290187675497882714112",
		"MaxInt16":	"int:32767",
		"MaxInt32":	"int:2147483647",
		"MaxInt64":	"int:9223372036854775807",
		"MaxInt8":	"int:127",
		"MaxUint16":	"int:65535",
		"MaxUint32":	"int:4294967295",
		"MaxUint64":	"int:18446744073709551615",
		"MaxUint8":	"int:255",
		"MinInt16":	"int:-32768",
		"MinInt32":	"int:-2147483648",
		"MinInt64":	"int:-9223372036854775808",
		"MinInt8":	"int:-128",
		"Phi":	"float:1355893059079843193030097015621153611996040418655586468942144325530191466769501318141688179946842310342863020505846766772662711907185210172534864277556413/837987995621412318723376562387865382967460363787024586107722590232610251879596686050117143635431464230626991136655378178359617675746660621652103062880256",
		"Pi":	"float:5265233861681329527430852685569923513156999664186156825699335632145130833652529926715379955630651875029455043336995646403337241248353639932329607400740841/1675975991242824637446753124775730765934920727574049172215445180465220503759193372100234287270862928461253982273310756356719235351493321243304206125760512",
		"SmallestNonzeroFloat32":	"float:13407807929942597099574024998205846127476859526281406305997157614053697519758116209723839542388419390837860564546892625342279307746235762181839259274038433/9568131466127621947163770315237577352582483950433331955534014747297500715462012198465648064079848065788579276806882658480425438483841942548911565191978080929321047135323978360596199778018349602045952",
		"SmallestNonzeroFloat64":	"float:837987995621412318723376562387865382967528542837591553054231272745319198424369655365935231952608328604793302723807080372339569516882534924689845329132395/169610658558250597457949882757253129256512582825232846524880306601538402823521241146029596619315984411186040958942601714115551210232864300981082979971362343290567722620469726473734686654453609615462992711162886914312002548467323965655079024463264900175579232518184090858573444982995938918601148906296146506082388046271954770533990741509051865409843001664221654092011338230906199419669398139261492786445466233786170361793273957887549496248290745722916589964061626271423952584704",
		"Sqrt2":	"float:2370187977027294181708131613913039815021135436007832465948110837130147201951018863264485267316490927583884794148767707476886477873843439514530955153130291/1675975991242824637446753124775730765934920727574049172215445180465220503759193372100234287270862928461253982273310756356719235351493321243304206125760512",
		"SqrtE":	"float:5526434531889553359100339527602551136535293759178014794854010166956579300433867493319128312730236202227030041787483535082815593803939658073242247126443667/3351951982485649274893506249551461531869841455148098344430890360930441007518386744200468574541725856922507964546621512713438470702986642486608412251521024",
		"SqrtPhi":	"float:4263748785949384222047394170590505028114438860640840609262555507082501186239921195173529474565705679454046425697050803084649101634749198923578942876946925/3351951982485649274893506249551461531869841455148098344430890360930441007518386744200468574541725856922507964546621512713438470702986642486608412251521024",
		"SqrtPi":	"float:5941180199407067869909325294831655792192555130773991214446196790352931403459697425080809383647149191660934688457292053536928055420572552259905515759885317/3351951982485649274893506249551461531869841455148098344430890360930441007518386744200468574541725856922507964546621512713438470702986642486608412251521024",
	},
	}
}
