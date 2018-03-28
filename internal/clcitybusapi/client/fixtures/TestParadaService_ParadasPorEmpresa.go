package fixtures

import (
	"encoding/json"
	"strconv"
	"testing"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

// TestParadaServiceParadasPorEmpresa fixture for test `TestParadaService_ParadasPorEmpresa`.
// Returns:
// 1- First `CodigoLineaParada` as int
// 2- First `CodigoLineaParada` as string
// 3- Second `CodigoLineaParada` as int
// 4- Second `CodigoLineaParada` as string
// 5- Lineas
// 6- Paradas by Linea
// 7- RecuperarLineasPorCodigoEmpresa request
// 8- RecuperarParadasCompletoPorLinea requests
// 9- RecuperarLineasPorCodigoEmpresa response
// 10- RecuperarParadasCompletoPorLinea responses
// 11- Expected output
func TestParadaServiceParadasPorEmpresa(t *testing.T) (
	l1 int,
	l1str string,
	l2 int,
	l2str string,
	fixl []*clcitybusapi.Linea,
	fixp map[string][]*clcitybusapi.Parada,
	flinreq *swparadas.RecuperarLineasPorCodigoEmpresa,
	fparreq [2]*swparadas.RecuperarParadasCompletoPorLinea,
	flinresp *swparadas.RecuperarLineasPorCodigoEmpresaResponse,
	fparresp [2]*swparadas.RecuperarParadasCompletoPorLineaResponse,
	fOut []*clcitybusapi.Parada,
) {
	l1 = 1529
	l1str = strconv.Itoa(l1)
	l2 = 1530
	l2str = strconv.Itoa(l2)

	fixl = []*clcitybusapi.Linea{
		&clcitybusapi.Linea{
			CodigoLineaParada: l1str,
			Descripcion:       "RAMAL A",
			CodigoEntidad:     "254",
			CodigoEmpresa:     355,
		},
		&clcitybusapi.Linea{
			CodigoLineaParada: l2str,
			Descripcion:       "RAMAL B",
			CodigoEntidad:     "254",
			CodigoEmpresa:     355,
		},
	}
	fixp = map[string][]*clcitybusapi.Parada{
		l1str: []*clcitybusapi.Parada{
			&clcitybusapi.Parada{
				Codigo:                     "57720",
				Identificador:              "RG001",
				Descripcion:                "HACIA CHACRA 11",
				AbreviaturaBandera:         "RAMAL A",
				AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
				LatitudParada:              "-53,803239",
				LongitudParada:             "-67,661785",
				AbreviaturaBanderaGIT:      "IDA A",
			},
			&clcitybusapi.Parada{
				Codigo:                     "57721",
				Identificador:              "RG002",
				Descripcion:                "HACIA CHACRA 11",
				AbreviaturaBandera:         "RAMAL A",
				AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
				LatitudParada:              "-53,803239",
				LongitudParada:             "-67,661785",
				AbreviaturaBanderaGIT:      "IDA A",
			},
		},
		l2str: []*clcitybusapi.Parada{
			&clcitybusapi.Parada{
				Codigo:                     "57725",
				Identificador:              "RG001",
				Descripcion:                "HACIA CHACRA Mi casa",
				AbreviaturaBandera:         "RAMAL B",
				AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
				LatitudParada:              "-53,803239",
				LongitudParada:             "-67,661785",
				AbreviaturaBanderaGIT:      "IDA B",
			},
			&clcitybusapi.Parada{
				Codigo:                     "57731",
				Identificador:              "RG003",
				Descripcion:                "HACIA asd 11",
				AbreviaturaBandera:         "RAMAL B",
				AbreviaturaAmpliadaBandera: "HACIA CHaaACRA 11",
				LatitudParada:              "-53,803239",
				LongitudParada:             "-67,661785",
				AbreviaturaBanderaGIT:      "IDA B",
			},
		},
	}

	// Fixture requests
	flinreq = &swparadas.RecuperarLineasPorCodigoEmpresa{
		Usuario:       "WEB.SUR",
		Clave:         "PAR.SW.SUR",
		CodigoEmpresa: 355,
		IsSublinea:    false,
	}
	fparreq = [2]*swparadas.RecuperarParadasCompletoPorLinea{
		&swparadas.RecuperarParadasCompletoPorLinea{
			Usuario:           "WEB.SUR",
			Clave:             "PAR.SW.SUR",
			CodigoLineaParada: int32(l1),
			IsSublinea:        false,
			IsInteligente:     false,
		},
		&swparadas.RecuperarParadasCompletoPorLinea{
			Usuario:           "WEB.SUR",
			Clave:             "PAR.SW.SUR",
			CodigoLineaParada: int32(l2),
			IsSublinea:        false,
			IsInteligente:     false,
		},
	}

	// Fixture results
	flinresu := &swparadas.RecuperarLineasPorCodigoEmpresaResult{
		CodigoEstado:  0,
		MensajeEstado: "ok",
		Lineas:        fixl,
	}
	flinresuJSON, err := json.Marshal(flinresu)
	if err != nil {
		t.Fatal("Error parsing JSON", err)
	}

	fparresu := [2]*swparadas.RecuperarParadasCompletoPorLineaResult{
		&swparadas.RecuperarParadasCompletoPorLineaResult{
			CodigoEstado:  0,
			MensajeEstado: "ok",
			Paradas:       fixp[l1str],
		},
		&swparadas.RecuperarParadasCompletoPorLineaResult{
			CodigoEstado:  0,
			MensajeEstado: "ok",
			Paradas:       fixp[l2str],
		},
	}

	var fparresuJSON [2][]byte
	for idx, result := range fparresu {
		resultJSON, err := json.Marshal(result)
		if err != nil {
			t.Fatal("Error parsing JSON", err)
		}
		fparresuJSON[idx] = resultJSON
	}

	// Fixture responses
	flinresp = &swparadas.RecuperarLineasPorCodigoEmpresaResponse{
		RecuperarLineasPorCodigoEmpresaResult: string(flinresuJSON),
	}

	fparresp = [2]*swparadas.RecuperarParadasCompletoPorLineaResponse{
		&swparadas.RecuperarParadasCompletoPorLineaResponse{
			RecuperarParadasCompletoPorLineaResult: string(fparresuJSON[0]),
		},
		&swparadas.RecuperarParadasCompletoPorLineaResponse{
			RecuperarParadasCompletoPorLineaResult: string(fparresuJSON[1]),
		},
	}

	// Fixture output
	fOut = append(fixp[l1str], fixp[l2str]...)

	return
}
