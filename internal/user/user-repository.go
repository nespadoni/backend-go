package user

import (
	"github.com/nespadoni/goerror"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	DB *gorm.DB
}

func NovoRepositorioUsuario(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{DB: db}
}

func (r *UsuarioRepository) BuscarTodos() []Usuario {
	var usuarios []Usuario
	resultado := r.DB.Find(&usuarios)
	goerror.EhErroBancoDados(resultado)
	return usuarios
}

func (r *UsuarioRepository) BuscarPorId(id int) Usuario {
	var usuario Usuario
	resultado := r.DB.Where("ID = ?", id).Find(&usuario)
	goerror.EhErroBancoDados(resultado)
	return usuario
}

func (r *UsuarioRepository) BuscarPorNome(nome string) Usuario {
	var usuario Usuario
	resultado := r.DB.Where("Nome = ?", nome).Find(&usuario)
	goerror.EhErroBancoDados(resultado)
	return usuario
}

func (r *UsuarioRepository) BuscarPorEmail(email string) Usuario {
	var usuario Usuario
	resultado := r.DB.Where("Email = ?", email).Find(&usuario)
	goerror.EhErroBancoDados(resultado)
	return usuario
}

func (r *UsuarioRepository) SalvarUsuario(usuario Usuario) Usuario {
	resultado := r.DB.Create(&usuario)
	goerror.EhErroBancoDados(resultado)
	return usuario
}

func (r *UsuarioRepository) AtualizarUsuario(id int, novoUsuario Usuario) {
	resultado := r.DB.Where("id = ?", id).Updates(novoUsuario)
	goerror.EhErroBancoDados(resultado)

}

func (r *UsuarioRepository) DeletarUsuario(id int) {
	usuario := Usuario{}
	resultado := r.DB.Where("id = ?", id).Delete(&usuario)
	goerror.EhErroBancoDados(resultado)
}
