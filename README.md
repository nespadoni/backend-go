# Rivaly - Sistema de Gestão Esportiva Universitária

![Go](https://img.shields.io/badge/Go-1.24-blue)
![Gin](https://img.shields.io/badge/Gin-Framework-green)
![GORM](https://img.shields.io/badge/GORM-ORM-orange)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-blue)
![Swagger](https://img.shields.io/badge/Swagger-Documentation-yellowgreen)

## 📋 Sobre o Projeto

O **Rivaly** é um aplicativo de gestão esportiva universitária voltado para atléticas universitárias, funcionando como redes sociais dedicadas ao esporte. A plataforma permite que:

- **Usuários comuns** acompanhem campeonatos, jogos, notícias e estatísticas
- **Usuários administradores** criem e gerenciem atléticas de suas universidades

## 🏗️ Arquitetura do Sistema

### Conceito Principal
O sistema funciona como uma rede social esportiva onde cada **Atlética** atua como um hub central que gerencia:
- Times
- Torneios e Campeonatos
- Notícias
- Jogadores
- Partidas e Estatísticas

### Hierarquia de Relacionamento
```

Usuário → Universidade → Atlética
    ↓
Atlética → Teams, Championships, News
    ↓
Championship → Tournaments → Matches → Results
    ↓
Teams → Players → Statistics
```
## 🔧 Tecnologias Utilizadas

- **Backend**: Go 1.24 com Gin Framework
- **ORM**: GORM
- **Banco de Dados**: PostgreSQL
- **Documentação**: Swagger/OpenAPI
- **Validação**: go-playground/validator
- **Configuração**: godotenv

## 🚀 Funcionalidades Principais

### 🏛️ Sistema de Permissões
O projeto implementa um sistema de permissões baseado em roles para cada atlética:

- **Owner**: Criador da atlética (controle total)
- **Admin**: Administrador com controle total
- **Moderator**: Permissões limitadas (criar notícias, mas não editar torneios)
- **Member**: Apenas visualiza e interage (curtir, comentar)

### 📊 Gestão de Dados Esportivos
- **Campeonatos** e **Torneios** organizados por modalidade esportiva
- **Partidas** com resultados em tempo real
- **Estatísticas detalhadas** por jogador (gols, assistências, cartões)
- **Escalações** e **lineups** para cada partida

### 🔔 Sistema de Notificações
Os usuários podem seguir diferentes atléticas e receber notificações sobre:
- Novos jogos e resultados
- Notícias importantes
- Atualizações de campeonatos

## 📁 Estrutura do Projeto
```

rivaly/
├── cmd/                    # Aplicação principal
├── config/                 # Configurações
├── docs/                   # Documentação Swagger
├── internal/               # Código interno da aplicação
│   ├── models/            # Modelos de dados
│   ├── user/              # Módulo de usuários
│   ├── championship/      # Módulo de campeonatos
│   └── athletic/          # Módulo de atléticas
├── pkg/                   # Pacotes reutilizáveis
│   ├── db/               # Configuração do banco
│   └── middleware/       # Middlewares
├── routes/               # Configuração de rotas
├── main.go              # Ponto de entrada
├── go.mod               # Dependências
└── README.md
```
## 🗄️ Modelo de Dados

### Entidades Principais

#### Usuários e Universidades
- `User` - Usuários do sistema
- `University` - Universidades cadastradas
- `Role` - Papéis/permissões

#### Atléticas e Times  
- `Athletic` - Atléticas universitárias
- `Team` - Times vinculados às atléticas
- `Player` - Jogadores dos times
- `UserRoleAthletic` - Relacionamento usuário-papel-atlética

#### Competições
- `Championship` - Campeonatos
- `Tournament` - Torneios dentro dos campeonatos
- `Sport` - Modalidades esportivas
- `Position` - Posições por modalidade

#### Jogos e Estatísticas
- `Match` - Partidas
- `Result` - Resultados das partidas
- `PlayerStats` - Estatísticas por jogador/partida
- `Lineup` - Escalações

#### Interações Sociais
- `News` - Notícias das atléticas
- `Follow` - Relacionamento usuário-atlética
- `Notification` - Sistema de notificações

## ⚙️ Configuração e Instalação

### Pré-requisitos
- Go 1.24+
- PostgreSQL
- Git

### 1. Clone o repositório
```
bash
git clone https://github.com/nespadoni/rivaly.git
cd rivaly
```
### 2. Configure as variáveis de ambiente
Crie um arquivo `.env` na raiz do projeto:
```
env
# Servidor
PORT=8080

# Banco de Dados
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=rivaly
```
### 3. Instale as dependências
```bash
go mod download
```
```


### 4. Execute as migrações
O sistema executa automaticamente as migrações na inicialização.

### 5. Execute a aplicação
```shell script
go run main.go
```


A aplicação estará disponível em `http://localhost:8080`

## 📚 Documentação da API

A documentação completa da API está disponível através do Swagger:

- **URL**: `http://localhost:8080/swagger/index.html`
- **Health Check**: `http://localhost:8080/health`

### Endpoints Principais

```
GET    /api/user          # Lista todos os usuários
GET    /api/user/:id      # Busca usuário por ID  
POST   /api/user          # Cria novo usuário
PUT    /api/user/:id      # Atualiza usuário
DELETE /api/user/:id      # Remove usuário

GET    /api/championship  # Lista campeonatos
```


## 🏃‍♂️ Como Executar

### Desenvolvimento
```shell script
# Execute em modo desenvolvimento
go run main.go

# Com hot reload (usando air)
air
```


### Produção
```shell script
# Build da aplicação
go build -o rivaly main.go

# Execute o binário
./rivaly
```


## 🧪 Testes

```shell script
# Execute todos os testes
go test ./...

# Testes com coverage
go test -cover ./...

# Testes verbose
go test -v ./...
```


## 🤝 Fluxo de Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add: nova funcionalidade'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

### Padrões de Commit
- `Add: nova funcionalidade`
- `Fix: correção de bug`
- `Update: atualização/melhoria`
- `Remove: remoção de código`
- `Docs: documentação`

## 📝 Roadmap

- [ ] Sistema de autenticação JWT
- [ ] Upload de imagens para times e atléticas
- [ ] Sistema de chat em tempo real
- [ ] Notificações push
- [ ] App mobile
- [ ] Dashboard analytics
- [ ] Sistema de ranking
- [ ] API de estatísticas avançadas

## 🐛 Issues Conhecidos

- Implementação de autenticação pendente
- Sistema de roles ainda em desenvolvimento
- Falta validação de permissões em alguns endpoints

## 👥 Equipe

- **Desenvolvedor Backend**: [Seu Nome]
- **Projeto**: Sistema de Gestão Esportiva Universitária

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 📞 Contato

- **Email**: contato@rivaly.com
- **GitHub**: [github.com/nespadoni/rivaly](https://github.com/nespadoni/rivaly)

---

⭐ **Se este projeto foi útil para você, considere dar uma estrela!**
