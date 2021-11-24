package keccak

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	qt "github.com/frankban/quicktest"
)

func TestKeccak(t *testing.T) {
	testKeccak(t, []byte("test"), "9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658")
	testKeccak(t, make([]byte, 32), "290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563")
	testKeccak(t, make([]byte, 100), "913fb9e1f6f1c6d910fd574a5cad8857aa43bfba24e401ada4f56090d4d997a7")
}

func testKeccak(t *testing.T, input []byte, expectedHex string) {
	expected := crypto.Keccak256(input)

	hBits := ComputeKeccak(bytesToBits(input))
	h := bitsToBytes(hBits)

	qt.Assert(t, h, qt.DeepEquals, expected)
	qt.Assert(t, hex.EncodeToString(h), qt.Equals, expectedHex)
}

func TestPad(t *testing.T) {
	b := make([]byte, 32)
	for i := 0; i < len(b); i++ {
		b[i] = byte(i)
	}
	bBits := bytesToBits(b)
	fBits := pad(bBits)

	qt.Assert(t, bitsToBytes(fBits[:]), qt.DeepEquals,
		[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
			16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128})
}

func TestKeccakfRound(t *testing.T) {
	s, _ := newS()

	s = keccakfRound(s, 0)
	qt.Assert(t, bitsToU64Array(s[:]), qt.DeepEquals,
		[]uint64{
			26388279066651, 246290629787648, 26388279902208,
			25165850, 246290605457408, 7784628352, 844424965783552,
			2305843009213694083, 844432714760192,
			2305843009249345539, 637534226, 14848, 641204224,
			14354, 3670528, 6308236288, 2130304761856,
			648518346341354496, 6309216256, 648520476645130240,
			4611706359392501763, 792677514882318336,
			20340965113972, 4611732197915754499,
			792633534417207412})

	s = keccakfRound(s, 20)
	qt.Assert(t, bitsToU64Array(s[:]), qt.DeepEquals,
		[]uint64{17728382861289829725, 13654073086381141005,
			9912591532945168756, 2030068283137172501, 5084683018496047808,
			151244976540463006, 11718217461613725815, 11636071286320763433,
			15039144509240642782, 11629028282864249197,
			2594633730779457624, 14005558505838459171, 4612881094252610438,
			2828009553220809993, 4838578484623267135, 1006588603063111352,
			11109191860075454495, 1187545859779038208,
			14661669042642437042, 5345317080454741069, 8196674451365552863,
			635818354583088260, 13515759754032305626, 1708499319988748543,
			7509292798507899312})

}

func TestKeccakf(t *testing.T) {
	s, _ := newS()

	s = keccakf(s)

	qt.Assert(t, bitsToU64Array(s[:]), qt.DeepEquals,
		[]uint64{9472389783892099349, 2159377575142921216,
			17826682512249813373, 2325963263767348549,
			15086930817298358378, 11661812091723830419,
			3517755057770134847, 5223775837645169598, 933274647126506074,
			3451250694486589320, 825065683101361807, 6192414258352188799,
			14426505790672879210, 3326742392640380689,
			16749975585634164134, 17847697619892908514,
			11598434253200954839, 6049795840392747215, 8610635351954084385,
			18234131770974529925, 15330347418010067760,
			12047099911907354591, 4763389569697138851, 6779624089296570504,
			15083668107635345971})

	// compute again keccakf on the current state
	s = keccakf(s)
	qt.Assert(t, bitsToU64Array(s[:]), qt.DeepEquals,
		[]uint64{269318771259381490, 15892848561416382510,
			12485559500958802382, 4360182510883008729,
			14284025675983944434, 8800366419087562177, 7881853509112258378,
			9503857914080778528, 17110477940977988953,
			13825318756568052601, 11460650932194163315,
			13272167288297399439, 13599957064256729412,
			12730838251751851758, 13736647180617564382,
			5651695613583298166, 15496251216716036782, 9748494184433838858,
			3637745438296580159, 3821184813198767406, 15603239432236101315,
			3726326332491237029, 7819962668913661099, 2285898735263816116,
			13518516210247555620})
}

func printBytes(name string, b []byte) {
	fmt.Printf("%s\n", name)
	for _, v := range b {
		fmt.Printf("\"%v\", ", v)
	}
	fmt.Println("")
}
func printU64Array(name string, b []uint64) {
	fmt.Printf("%s\n", name)
	for _, v := range b {
		fmt.Printf("\"%v\", ", v)
	}
	fmt.Println("")
}

func TestAbsorb(t *testing.T) {
	s, _ := newS()
	block := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128}
	// printU64Array("s", bitsToU64Array(s[:]))
	// printBytes("block", block[:])

	absorbed := absorb(s, bytesToBits(block))
	// printU64Array("absorbed", bitsToU64Array(absorbed[:]))

	qt.Assert(t, bitsToU64Array(absorbed[:]), qt.DeepEquals,
		[]uint64{8342348566319207042, 319359607942176202, 14410076088654599075,
			15666111399434436772, 9558421567405313402, 3396178318116504023,
			794353847439963108, 12717011319735989377, 3503398863218919239,
			5517201702366862678, 15999361614129160496, 1325524015888689985,
			11971708408118944333, 14874486179441062217, 12554876384974234666,
			11129975558302206043, 11257826431949606534, 2740710607956478714,
			15000019752453010167, 15593606854132419294, 2598425978562809333,
			8872504799797239246, 1212062965004664308, 5443427421087086722,
			10946808592826700411})

	absorbed = absorb(absorbed, bytesToBits(block))
	// printU64Array("absorbed", bitsToU64Array(absorbed[:]))

	qt.Assert(t, bitsToU64Array(absorbed[:]), qt.DeepEquals,
		[]uint64{8909243822027471379, 1111840847970088140,
			12093072708540612559, 11255033638786021658, 2082116894939842214,
			12821085060245261575, 6901785969834988344, 3182430130277914993,
			2164708585929408975, 14402143231999718904, 16231444410553803968,
			1850945423480060493, 12856855675247400303, 1137248620532111171,
			7389129221921446308, 12932467982741614601, 1350606937385760406,
			10983682292859713641, 10305595434820307765, 13958651111365489854,
			17206620388135196198, 4238113785249530092, 7230868147643218103,
			603011106238724524, 16480095441097880488})
}

func TestFinal(t *testing.T) {
	b := make([]byte, 32)
	for i := 0; i < len(b); i++ {
		b[i] = byte(i)
	}
	bBits := bytesToBits(b)

	fBits := final(bBits)

	// printBytes("in", b[:])
	// printU64Array("out", bitsToU64Array(fBits[:]))

	qt.Assert(t, bitsToU64Array(fBits[:]), qt.DeepEquals,
		[]uint64{16953415415620100490, 7495738965189503699,
			12723370805759944158, 3295955328722933810,
			12121371508560456016, 174876831679863147, 15944933357501475584,
			7502339663607726274, 12048918224562833898,
			16715284461100269102, 15582559130083209842,
			1743886467337678829, 2424196198791253761, 1116417308245482383,
			10367365997906434042, 1849801549382613906,
			13294939539683415102, 4478091053375708790, 2969967870313332958,
			14618962068930014237, 2721742233407503451,
			12003265593030191290, 8109318293656735684, 6346795302983965746,
			12210038122000333046})

	// 2nd test

	for i := 0; i < len(b); i++ {
		b[i] = byte(254)
	}
	bBits = bytesToBits(b)
	fBits = final(bBits)

	// printBytes("in", b[:])
	// printU64Array("out", bitsToU64Array(fBits[:]))
	qt.Assert(t, bitsToU64Array(fBits[:]), qt.DeepEquals,
		[]uint64{16852464862333879129, 9588646233186836430, 693207875935078627,
			6545910230963382296, 3599194178366828471, 13130606490077331384,
			10374798023615518933, 7285576075118720444, 4097382401500492461,
			3968685317688314807, 3350659309646210303, 640023485234837464,
			2550030127986774041, 8948768022010378840, 10678227883444996205,
			1395278318096830339, 2744077813166753978, 13362598477502046010,
			14601579319881128511, 4070707967569603186, 16833768365875755098,
			1486295134719870048, 9161068934282437999, 8245604251371175619,
			8421994351908003183})
}

func TestSqueeze(t *testing.T) {
	in := []uint64{16852464862333879129, 9588646233186836430, 693207875935078627,
		6545910230963382296, 3599194178366828471, 13130606490077331384,
		10374798023615518933, 7285576075118720444, 4097382401500492461,
		3968685317688314807, 3350659309646210303, 640023485234837464,
		2550030127986774041, 8948768022010378840, 10678227883444996205,
		1395278318096830339, 2744077813166753978, 13362598477502046010,
		14601579319881128511, 4070707967569603186, 16833768365875755098,
		1486295134719870048, 9161068934282437999, 8245604251371175619,
		8421994351908003183}
	inBits := u64ArrayToBits(in)
	var inBits1600 [25 * 64]bool
	copy(inBits1600[:], inBits[:])

	outBits := squeeze(inBits1600)

	// printU64Array("in", in)
	// printBytes("out", bitsToBytes(outBits[:]))

	qt.Assert(t, bitsToBytes(outBits[:]), qt.DeepEquals,
		[]byte{89, 195, 41, 13, 129, 251, 223, 233, 206, 31, 253, 61,
			242, 182, 17, 133, 227, 8, 157, 240, 227, 196, 158, 9, 24, 232,
			42, 96, 172, 190, 215, 90})

	// 2nd test
	in = []uint64{16953415415620100490, 7495738965189503699,
		12723370805759944158, 3295955328722933810, 12121371508560456016,
		174876831679863147, 15944933357501475584, 7502339663607726274,
		12048918224562833898, 16715284461100269102, 15582559130083209842,
		1743886467337678829, 2424196198791253761, 1116417308245482383,
		10367365997906434042, 1849801549382613906, 13294939539683415102,
		4478091053375708790, 2969967870313332958, 14618962068930014237,
		2721742233407503451, 12003265593030191290, 8109318293656735684,
		6346795302983965746, 12210038122000333046}
	inBits = u64ArrayToBits(in)
	copy(inBits1600[:], inBits[:])

	outBits = squeeze(inBits1600)

	// printU64Array("in", in)
	// printBytes("out", bitsToBytes(outBits[:]))
	qt.Assert(t, bitsToBytes(outBits[:]), qt.DeepEquals,
		[]byte{138, 225, 170, 89, 127, 161, 70, 235, 211, 170, 44, 237,
			223, 54, 6, 104, 222, 165, 229, 38, 86, 126, 146, 176, 50, 24,
			22, 164, 232, 149, 189, 45})
}
