package otpx

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/hotp"
	"github.com/pquerna/otp/totp"
	"image/png"
	"math"
	"os"
	"time"
)

const SEED = "xwqjdi"



type Auth struct {
	Code string
	Path string
}

func GenOtpKey(issuer string, account string) (*Auth, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: account,
		Period:      30,
		SecretSize:  10,
		Secret:      []byte{},
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
		Rand:        rand.Reader,
	})
	if err != nil {
		return nil, err
	}

	img, err := key.Image(200, 200)
	if err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%x", md5.Sum([]byte(issuer+account+SEED))) + ".png"
	path := "qrcode/" + fileName
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	return &Auth{
		Code: key.Secret(),
		Path: path,
	}, nil
}

func Validate(code string, secret string) bool {
	//res, _ := totp.ValidateCustom(code, secret,
	//	time.Now().UTC(),
	//	totp.ValidateOpts{
	//		Period:    30,
	//		Skew:      15,
	//		Digits:    otp.DigitsSix,
	//		Algorithm: otp.AlgorithmSHA1,
	//	})
	counter := uint64(math.Floor(float64(time.Now().UTC().Unix()) / float64(30)))
	res := hotp.Validate(code, counter, secret)
	return res
}
