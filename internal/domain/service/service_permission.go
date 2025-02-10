//service permission
package service
import(
 "api/internal/domain/models"  
 "api/internal/database"
)
type(
    Permission = models.Permission
)
func GetPermission() []Permission{
    var permission []Permission
    database.DB.Find(&permission)
    return permission
}
func CreatePermission(permission Permission) Permission{
    database.DB.Create(&permission)
    return permission
}
func FindPermission(id uint) Permission{
    var permission Permission
    database.DB.Take(id, &permission)
    return permission
}
func UpdatePermission(id uitn,permissions Permission) Permission{
   permission := Permission{
       Name : permissions.Name,
   }
   database.DB.Save(&permission)
   return permission
}
func DeletePermission(id uint)(string,error){
    var permission Permission
    if err := database.DB.Take(&permission,id).Error;err != nil{
        return "",err
    }
    database.DB.Delete(&permission)
    return "Permission success deleted"
}