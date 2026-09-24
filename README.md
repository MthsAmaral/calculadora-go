# Calculadora em Go — Servidor HTTP

Aplicação web simples para atividade de implantação: calculadora (soma, subtração, multiplicação e divisão) e exibição do horário atual da instância Ubuntu. O servidor HTTP nativo de Go responde na porta 3000.

## Arquivos

- `main.go`: aplicação inteira, incluindo servidor HTTP, lógica da calculadora e HTML/CSS.
- `calculadora-go.service`: arquivo de serviço systemd, para manter o programa em execução.

## Instalar Go

```bash
sudo apt update
sudo apt install -y golang-go
```

## Testar em desenvolvimento

Dentro da pasta do projeto:

```bash
go run main.go
```

Em uma segunda sessão SSH:

```bash
curl http://localhost:3000
```

Teste uma conta via curl:

```bash
curl -s -X POST \
  -d "numero1=10" \
  -d "numero2=5" \
  -d "operacao=soma" \
  http://localhost:3000 | grep "Resultado"
```

## Compilar

```bash
go build -o calculadora-go main.go
./calculadora-go
```

## Configurar serviço systemd

Pare a aplicação manual com Ctrl+C antes de criar o serviço.

```bash
sudo mkdir -p /opt/calculadora-go
sudo cp calculadora-go /opt/calculadora-go/
sudo chmod 755 /opt/calculadora-go/calculadora-go
sudo cp calculadora-go.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now calculadora-go
```

Confira o serviço:

```bash
sudo systemctl status calculadora-go --no-pager
sudo systemctl is-enabled calculadora-go
```

Teste novamente:

```bash
curl http://localhost:3000
```

## Acesso externo

Libere a porta 3000 no firewall local, se o UFW estiver ativo:

```bash
sudo ufw allow 3000/tcp
```

Também é necessário liberar TCP/3000 na regra de entrada da sua instância/provedor. Acesse usando:

```text
http://IP_PUBLICO_DA_INSTANCIA:3000
```
