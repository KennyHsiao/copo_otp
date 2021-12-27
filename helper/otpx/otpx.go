package otpx

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"image/png"
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
		SecretSize:  20,
		Secret:      []byte{},
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA256,
		Rand:        rand.Reader,
	})
	if err != nil {
		return nil, err
	}

	img, err := key.Image(200, 200)
	if err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%x", md5.Sum([]byte(issuer+account+SEED))) +".png"
	path := "qrcode/"+fileName
	f, err := os.Create("../"+path)
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

func Validate(code string, secret string) (bool) {
	res, _ :=totp.ValidateCustom(code, secret,
		time.Now().UTC(),
		totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA256,
	})

	return res
}