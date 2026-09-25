package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type PaginaDados struct {
	Horario   string
	Resultado string
	Erro      string
	Numero1   string
	Numero2   string
	Operacao string
}

const paginaHTML = `
<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Calculadora Go - HelloWorld</title>
    <style>
        * { box-sizing: border-box; }
        body {
            margin: 0; min-height: 100vh; padding: 20px; display: flex;
            align-items: center; justify-content: center;
            font-family: Arial, Helvetica, sans-serif;
            background: linear-gradient(135deg, #1e3c72, #2a5298);
        }
        .card {
            width: 100%; max-width: 430px; padding: 32px; border-radius: 12px;
            background: #fff; box-shadow: 0 10px 30px rgba(0, 0, 0, .30);
        }
        h1 { margin: 0 0 6px; color: #1e3c72; text-align: center; font-size: 24px; }
        .subtitulo { margin-bottom: 22px; color: #666; text-align: center; font-size: 14px; }
        .relogio {
            margin-bottom: 20px; padding: 12px; border-radius: 8px;
            background: #1e3c72; color: #fff; text-align: center; font-size: 15px;
        }
        label { display: block; margin-top: 14px; color: #333; font-size: 14px; font-weight: bold; }
        input, select, button { width: 100%; margin-top: 6px; padding: 11px; border-radius: 6px; font-size: 16px; }
        input, select { border: 1px solid #c9c9c9; background: #fff; }
        button { margin-top: 22px; border: 0; background: #2a5298; color: #fff; cursor: pointer; font-weight: bold; }
        button:hover { background: #1e3c72; }
        .resultado {
            margin-top: 20px; padding: 14px; border-left: 4px solid #2a5298;
            border-radius: 6px; background: #eef3fb; color: #1e3c72; text-align: center;
        }
        .erro {
            margin-top: 20px; padding: 14px; border-left: 4px solid #d93025;
            border-radius: 6px; background: #fdecea; color: #b42318; text-align: center;
        }
        .rodape { margin-top: 22px; color: #888; text-align: center; font-size: 12px; }
    </style>
</head>
<body>
    <main class="card">
        <h1>Calculadora em Go</h1>
        <div class="subtitulo">Aplicação HelloWorld — Servidor HTTP</div>

        <div class="relogio">Horário atual do servidor: {{.Horario}}</div>

        <form method="post" action="/">
            <label for="numero1">Primeiro número</label>
            <input id="numero1" type="text" name="numero1" value="{{.Numero1}}" placeholder="Ex.: 10" required>

            <label for="operacao">Operação</label>
            <select id="operacao" name="operacao">
                <option value="soma" {{if eq .Operacao "soma"}}selected{{end}}>Soma (+)</option>
                <option value="subtracao" {{if eq .Operacao "subtracao"}}selected{{end}}>Subtração (-)</option>
                <option value="multiplicacao" {{if eq .Operacao "multiplicacao"}}selected{{end}}>Multiplicação (*)</option>
                <option value="divisao" {{if eq .Operacao "divisao"}}selected{{end}}>Divisão (/)</option>
            </select>

            <label for="numero2">Segundo número</label>
            <input id="numero2" type="text" name="numero2" value="{{.Numero2}}" placeholder="Ex.: 5" required>

            <button type="submit">Calcular</button>
        </form>

        {{if .Resultado}}
        <div class="resultado">Resultado: <strong>{{.Resultado}}</strong></div>
        {{end}}

        {{if .Erro}}
        <div class="erro">{{.Erro}}</div>
        {{end}}

        <div class="rodape">Ubuntu + Go + systemd</div>
    </main>
</body>
</html>
`

var pagina = template.Must(template.New("calculadora").Parse(paginaHTML))

func calculadoraHandler(w http.ResponseWriter, r *http.Request) {
	dados := PaginaDados{
		Horario:   time.Now().Format("02/01/2006 15:04:05"),
		Operacao: "soma",
	}

	if r.Method == http.MethodPost {
		dados.Numero1 = r.FormValue("numero1")
		dados.Numero2 = r.FormValue("numero2")
		dados.Operacao = r.FormValue("operacao")

		n1, erro1 := strconv.ParseFloat(dados.Numero1, 64)
		n2, erro2 := strconv.ParseFloat(dados.Numero2, 64)

		if erro1 != nil || erro2 != nil {
			dados.Erro = "Digite dois números válidos."
		} else {
			var resultado float64

			switch dados.Operacao {
			case "soma":
				resultado = n1 + n2
			case "subtracao":
				resultado = n1 - n2
			case "multiplicacao":
				resultado = n1 * n2
			case "divisao":
				if n2 == 0 {
					dados.Erro = "Não é possível dividir por zero."
				} else {
					resultado = n1 / n2
				}
			default:
				dados.Erro = "Operação inválida."
			}

			if dados.Erro == "" {
				dados.Resultado = fmt.Sprintf("%.2f", resultado)
			}
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := pagina.Execute(w, dados); err != nil {
		http.Error(w, "Erro ao gerar a página.", http.StatusInternalServerError)
		log.Println("Erro no template:", err)
	}
}

func main() {
	http.HandleFunc("/", calculadoraHandler)

	porta := ":3001"
	log.Printf("Servidor HTTP iniciado em http://localhost%s", porta)
	log.Fatal(http.ListenAndServe(porta, nil))
}
