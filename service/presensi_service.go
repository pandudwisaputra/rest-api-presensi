package service

import (
	"golang.org/x/net/context"
	"golang/rest-api-presensi/entity/web"
)

type PresensiService interface {
	PresensiMasuk(ctx context.Context, request web.PresensiMasukRequest) web.PresensiMasukResponse
	PresensiTidakMasuk(ctx context.Context, request web.PresensiTidakMasukRequest) web.PresensiTidakMasukResponse
	PresensiKeluar(ctx context.Context, request web.PresensiKeluarRequest) web.PresensiKeluarResponse
	Riwayat(ctx context.Context, request int) []web.RiwayatPresensiResponse
	PresensiCheck(ctx context.Context, request int) web.PresensiCheckResponse
}
