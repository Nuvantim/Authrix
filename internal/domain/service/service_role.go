//role service
package service
import(
    "api/internal/models"
    "api/internal/database"
)
type(
    var Role = models.Role
)
func GetRole()[]Role{
    var role []Role
    database.DB.Find(&role)
    return role
}
func CreateRole()Role{}
func FindRole(id uint)Role{
    var role Role
    database.DB.Preload("Permission").Take(&role,id)
    return role
}
func UpdateRole()Role{}
func DeleteRole() error{
var role models.Role
    if err := database.DB.Take(&role, id).Error; err != nil {
        return err
    }

    // Hapus relasi dari pivot table
    database.DB.Model(&role).Association("Permissions").Clear()

    // Hapus role dari database
    database.DB.Delete(&role)

    return nil
}