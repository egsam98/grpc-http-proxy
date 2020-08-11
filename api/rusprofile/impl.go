package rusprofile

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const RusprofileHost = "rusprofile.ru"

type RusprofileServer struct {
	logger grpclog.LoggerV2
}

func NewRusprofileServer(logger grpclog.LoggerV2) *RusprofileServer {
	return &RusprofileServer{logger}
}

func (rs *RusprofileServer) FetchCounterparty(_ context.Context, inn *INN) (*Counterparty, error) {
	if !inn.isValid() {
		return nil, status.Error(codes.InvalidArgument, "некорректно введен ИНН, проверьте кол-во цифр")
	}

	res, err := http.Get(fmt.Sprintf("https://%s/search?query=%d", RusprofileHost, inn.Id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		if res.StatusCode == http.StatusTooManyRequests {
			return nil, status.Error(codes.ResourceExhausted, "превышено число обращений к "+RusprofileHost)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			rs.logger.Error(err)
			return nil, err
		}
		rs.logger.Error(string(body))
		code := codes.InvalidArgument
		if res.StatusCode >= 500 {
			code = codes.Unavailable
		}
		return nil, status.Error(code, RusprofileHost+" "+res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	resultHeader := doc.Find(".company-header__row > h1").Text()
	if strings.HasSuffix(resultHeader, "0 организаций и 0 индивидуальных предпринимателей") {
		return nil, status.Error(codes.NotFound, resultHeader)
	}

	kpp, err := strconv.Atoi(doc.Find("#clip_kpp").Text())
	if err != nil {
		return nil, err
	}

	return &Counterparty{
		INN:   inn.Id,
		KPP:   uint64(kpp),
		Title: doc.Find(".company-name").Text(),
		Head:  doc.Find(".gtm_main_fl > span").Text(),
	}, nil
}

func (inn *INN) isValid() bool {
	count := 0
	number := inn.Id
	for number != 0 {
		number /= 10
		count += 1
	}
	return count == 12 || count == 10
}
