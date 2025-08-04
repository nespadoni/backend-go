package user

//
//import (
//	"gorm.io/gorm"
//)
//
//type UsuarioRepository struct {
//	DB *gorm.DB
//}
//
//func NovoRepositorioUsuario(db *gorm.DB) *UsuarioRepository {
//	return &UsuarioRepository{DB: db}
//}
//
//func (r *UsuarioRepository) BuscarTodos() []Usuario {
//	var usuarios []Usuario
//	resultado := r.DB.Find(&usuarios)
//	handler.NewReadError("Erro ao buscar todos os usuários: ", resultado.Error)
//	return usuarios
//}
//
//func (r *UsuarioRepository) BuscarPorId(id int) Usuario {
//	var usuario Usuario
//	resultado := r.DB.Where("ID = ?", id).Find(&usuario)
//	handler.NewReadError("Erro ao buscar usuário por ID", resultado.Error)
//	return usuario
//}
//
//func (r *UsuarioRepository) BuscarPorNome(nome string) Usuario {
//	var usuario Usuario
//	resultado := r.DB.Where("Nome = ?", nome).Find(&usuario)
//	handler.NewReadError("Erro ao buscar usuário por nome", resultado.Error)
//	return usuario
//}
//
//func (r *UsuarioRepository) BuscarPorEmail(email string) Usuario {
//	var usuario Usuario
//	resultado := r.DB.Where("Email = ?", email).Find(&usuario)
//	handler.Novo("Erro ao buscar usuário por email", resultado.Error)
//	return usuario
//}
//
//func (r *UsuarioRepository) SalvarUsuario(usuario Usuario) Usuario {
//	resultado := r.DB.Create(&usuario)
//	handler.NewCreateError("Erro ao salvar usuário: ", resultado.Error)
//	return usuario
//}
//
//func (r *UsuarioRepository) AtualizarUsuario(id int, novoUsuario Usuario) {
//	resultado := r.DB.Where("id = ?", id).Updates(novoUsuario)
//	handler.NewUpdateError("Erro ao atualizar usuário: ", resultado.Error)
//}
//
//func (r *UsuarioRepository) DeletarUsuario(id int) {
//	usuario := Usuario{}
//	resultado := r.DB.Where("id = ?", id).Delete(&usuario)
//	handler.NewDeleteError("Erro ao deletar usuário: ", resultado.Error)
//}
