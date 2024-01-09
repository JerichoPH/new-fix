package controllers

import (
	"fmt"
	"new-fix/models"
	"new-fix/utils"
	"new-fix/wrongs"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type (
	// AccountCtrl 用户控制器
	AccountCtrl struct{}
	// AccountStoreForm 新建用户表单
	AccountStoreForm struct {
		Username             string   `json:"username"`
		Nickname             string   `json:"nickname"`
		Password             string   `json:"password"`
		PasswordConfirmation string   `json:"password_confirmation"`
		RbacRoleUuids        []string `json:"rbac_role_uuids"`
		Avatar               struct {
			OriginalFilename  string  `json:"original_filename"`
			OriginalExtension string  `json:"original_extension"`
			Base64            *string `json:"base64"`
		} `json:"avatar"`
		rbacRoles []*models.RbacRoleMdl
	}
	// AccountUpdateForm 编辑用户表单
	AccountUpdateForm struct {
		Username      string   `json:"username"`
		Nickname      string   `json:"nickname"`
		RbacRoleUuids []string `json:"rbac_role_uuids"`
		Avatar        struct {
			OriginalFilename  string  `json:"original_filename"`
			OriginalExtension string  `json:"original_extension"`
			Base64            *string `json:"base64"`
		} `json:"avatar"`
		rbacRoles []*models.RbacRoleMdl
	}
	// AccountUpdatePasswordForm 修改密码表单
	AccountUpdatePasswordForm struct {
		OldPassword          string `json:"old_password"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
	}
	// AccountDestroyManyForm 批量删除
	AccountDestroyManyForm struct {
		Uuids    []string `json:"uuids"`
		accounts []*models.AccountMdl
	}
)

// NewAccountCtrl 构造函数
func NewAccountCtrl() *AccountCtrl {
	return &AccountCtrl{}
}

// ShouldBind 新建用户表单绑定
func (receiver AccountStoreForm) ShouldBind(ctx *gin.Context) AccountStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate(err.Error())
	}
	if len(receiver.Username) < 2 {
		wrongs.ThrowValidate("账号不能小于2位")
	}
	if len(receiver.Nickname) < 2 {
		wrongs.ThrowValidate("昵称不能小于2位")
	}
	if len(receiver.Password) < 6 {
		wrongs.ThrowValidate("密码不能小于6位")
	}
	if receiver.Password != receiver.PasswordConfirmation {
		wrongs.ThrowValidate("两次密码不一致")
	}
	if len(receiver.RbacRoleUuids) > 0 {
		models.NewRbacRoleMdl().GetDb("").Where("uuid in ?", receiver.RbacRoleUuids).Find(&receiver.rbacRoles)
	}

	return receiver
}

// ShouldBind 编辑用户表单绑定
func (receiver AccountUpdateForm) ShouldBind(ctx *gin.Context) AccountUpdateForm {
	if ctx.Param("uuid") == "" {
		wrongs.ThrowValidate("用户编号不能为空")
	}
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate(err.Error())
	}
	if len(receiver.Username) < 2 {
		wrongs.ThrowValidate("账号不能小于2位")
	}
	if len(receiver.Nickname) < 2 {
		wrongs.ThrowValidate("昵称不能小于2位")
	}
	if len(receiver.RbacRoleUuids) > 0 {
		models.NewRbacRoleMdl().GetDb("").Where("uuid in ?", receiver.RbacRoleUuids).Find(&receiver.rbacRoles)
	}
	if receiver.Avatar.Base64 != nil {
		if receiver.Avatar.OriginalFilename == "" {
			wrongs.ThrowValidate("文件名不能为空")
		}
		if receiver.Avatar.OriginalExtension == "" {
			wrongs.ThrowValidate("文件扩展名不能为空")
		}
	}

	return receiver
}

// ShouldBind 修改用户密码表单绑定
func (receiver AccountUpdatePasswordForm) ShouldBind(ctx *gin.Context) AccountUpdatePasswordForm {
	if ctx.Param("uuid") == "" {
		wrongs.ThrowValidate("用户编号不能为空")
	}

	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate(err.Error())
	}
	if len(receiver.OldPassword) < 6 {
		wrongs.ThrowValidate("原密码不能小于6位")
	}
	if len(receiver.Password) < 6 {
		wrongs.ThrowValidate("新密码不能小于6位")
	}
	if receiver.Password != receiver.PasswordConfirmation {
		wrongs.ThrowValidate("两次密码不一致")
	}

	return receiver
}

// ShouldBind 批量删除用户表单绑定
func (receiver AccountDestroyManyForm) ShouldBind(ctx *gin.Context) AccountDestroyManyForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.ThrowValidate("表单绑定错误：%s", err.Error())
	}

	if len(receiver.Uuids) > 0 {
		models.NewAccountMdl().GetDb("").Where("uuid in ?", receiver.Uuids).Find(&receiver.accounts)
	}

	return receiver
}

// Store 新建
func (AccountCtrl) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		err    error
		repeat *models.AccountMdl
		avatar *models.FileMdl
	)

	// 表单绑定
	form := AccountStoreForm{}.ShouldBind(ctx)

	// 查重
	ret = models.NewAccountMdl().GetDb("").Where("username = ?", form.Username).First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "账号")
	ret = models.NewAccountMdl().GetDb("").Where("nickname = ?", form.Nickname).First(&repeat)
	wrongs.ThrowWhenNotEmpty(ret, "昵称")

	// 新建
	account := &models.AccountMdl{
		MySqlMdl: models.MySqlMdl{Uuid: uuid.NewV4().String()},
		Username: form.Username,
		Nickname: form.Nickname,
		Password: utils.GeneratePassword(form.Password),
		BeAdmin:  false,
	}
	if ret = models.
		NewAccountMdl().
		GetDb("").
		Create(&account); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	// 保存头像
	if form.Avatar.Base64 != nil {
		var savePath, base64 string

		savePath = fmt.Sprintf("public/avatars/%s", account.Uuid)
		base64 = *form.Avatar.Base64
		if avatar, err = (&models.FileMdl{}).CreateByBase64(
			savePath,
			form.Avatar.OriginalFilename,
			form.Avatar.OriginalExtension,
			base64,
		); err != nil {
			wrongs.ThrowForbidden("保存用户头像失败：%s", err.Error())
		} else {
			account.AvatarUuid = avatar.Uuid
			if ret = models.
				NewAccountMdl().
				GetDb("").
				Where("uuid = ?", account.Uuid).
				Save(&account); ret.Error != nil {
				wrongs.ThrowForbidden("保存用户头像失败：%s", ret.Error.Error())
			}
		}
	}

	// 用户绑定角色
	account.BindRbacRoles(form.rbacRoles)

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Created(map[string]any{"account": account}).ToGinResponse())
}

// Destroy 删除
func (AccountCtrl) Destroy(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		account models.AccountMdl
	)

	// 查询
	ret = models.
		NewAccountMdl().
		GetDb("").
		Where("uuid = ?", ctx.Param("uuid")).
		First(&account)
	wrongs.ThrowWhenEmpty(ret, "用户")

	// 删除
	if ret := models.NewAccountMdl().GetDb("").Where("uuid = ?", account.Uuid).Delete(&account); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// DestroyMany 批量删除
func (AccountCtrl) DestroyMany(ctx *gin.Context) {
	form := AccountDestroyManyForm{}.ShouldBind(ctx)

	if len(form.accounts) > 0 {
		models.NewAccountMdl().GetDb("").Delete(&form.accounts)
	}

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (AccountCtrl) Update(ctx *gin.Context) {
	var (
		ret           *gorm.DB
		err           error
		account       *models.AccountMdl
		avatar        *models.FileMdl
		oldAvatarUuid string
		// repeat  *models.AccountModel
	)

	// 表单
	form := new(AccountUpdateForm).ShouldBind(ctx)

	// 查重
	ret = models.NewAccountMdl().GetDb("").Where("username = ?", form.Username).Where("uuid <> ?", ctx.Param("uuid")).First(nil)
	wrongs.ThrowWhenNotEmpty(ret, "账号被占用")
	ret = models.NewAccountMdl().GetDb("").Where("nickname = ?", form.Nickname).Where("uuid <> ?", ctx.Param("uuid")).First(nil)
	wrongs.ThrowWhenNotEmpty(ret, "昵称被占用")

	// 查询
	ret = models.NewAccountMdl().GetDb("").Where("uuid", ctx.Param("uuid")).First(&account)
	wrongs.ThrowWhenEmpty(ret, "用户")

	// 编辑
	account.Username = form.Username
	account.Nickname = form.Nickname

	// 保存修改
	if ret = models.NewAccountMdl().GetDb("").Where("uuid = ?", ctx.Param("uuid")).Save(&account); ret.Error != nil {
		wrongs.ThrowForbidden(ret.Error.Error())
	}

	// 保存用户头像
	oldAvatarUuid = account.AvatarUuid
	if form.Avatar.Base64 != nil {
		var savePath, base64 string

		savePath = fmt.Sprintf("public/avatars/%s", account.Uuid)
		base64 = *form.Avatar.Base64
		if avatar, err = (&models.FileMdl{}).CreateByBase64(
			savePath,
			form.Avatar.OriginalFilename,
			form.Avatar.OriginalExtension,
			base64,
		); err != nil {
			wrongs.ThrowForbidden("保存用户头像失败：%s", err.Error())
		} else {
			account.AvatarUuid = avatar.Uuid
			if ret = models.
				NewAccountMdl().
				GetDb("").
				Where("uuid = ?", account.Uuid).
				Save(&account); ret.Error != nil {
				wrongs.ThrowForbidden("保存用户头像失败：%s", ret.Error.Error())
			}
			// 删除旧头像
			if err = (&models.FileMdl{}).DeleteFile(oldAvatarUuid); err != nil {
				wrongs.ThrowForbidden("删除旧头像失败：%s", err.Error())
			}
		}
	}

	// 用户绑定角色
	account.BindRbacRoles(form.rbacRoles)

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Updated(map[string]any{"account": account}).ToGinResponse())
}

// PostUpdateAvatar 更新用户头像
func (AccountCtrl) PostUpdateAvatar(ctx *gin.Context) {
	var (
		ret                  *gorm.DB
		newAvatar, oldAvatar *models.FileMdl
		account              *models.AccountMdl
	)

	file, err := ctx.FormFile("avatar")
	if err != nil {
		wrongs.ThrowForbidden("上传文件失败：%s", err.Error())
	}

	accountUuid := ctx.Param("uuid")

	ret = models.NewAccountMdl().GetDb("").Where("uuid = ?", accountUuid).First(&account)
	wrongs.ThrowWhenEmpty(ret, "用户")

	// 你可以改变保存的文件名
	relativePath := path.Join("public", "avatars", account.Uuid)
	fs := utils.NewFileSystem(utils.GetRootPath()).Join([]string{relativePath})
	if !fs.IsExist() {
		fs.CreateDir()
	}

	newAvatarUuid := uuid.NewV4().String()
	originalFilename := file.Filename
	originalExection := strings.Split(originalFilename, ".")[len(strings.Split(originalFilename, "."))-1]
	saveFilename := fmt.Sprintf("%s.%s", newAvatarUuid, originalExection)

	// 创建文件对象
	newAvatar = &models.FileMdl{
		MySqlMdl:          models.MySqlMdl{Uuid: newAvatarUuid},
		SavePath:          relativePath,
		Filename:          saveFilename,
		OriginalFilename:  originalFilename,
		OriginalExtension: originalExection,
	}
	if ret = models.NewFileMdl().GetDb("").Create(&newAvatar); ret.Error != nil {
		wrongs.ThrowForbidden("保存头像文件数据失败：%s", ret.Error.Error())
	}

	if err := ctx.SaveUploadedFile(file, path.Join(fs.GetPath(), newAvatar.Filename)); err != nil {
		wrongs.ThrowForbidden("保存头像文件失败：%s", err.Error())
	}

	// 修改头像
	oldAvatarUuid := account.AvatarUuid
	if oldAvatarUuid != "" {
		// 删除旧头像
		if err = oldAvatar.DeleteFile(oldAvatarUuid); err != nil {
			wrongs.ThrowForbidden("删除旧头像文件失败：%s", err.Error())
		}
	}

	// 修改用户头像
	models.NewAccountMdl().GetDb("").Where("uuid = ?", accountUuid).Update("avatar_uuid", newAvatar.Uuid)

	// 重读新头像文件数据
	models.NewFileMdl().GetDb("").Where("uuid = ?", newAvatarUuid).First(&newAvatar)

	ctx.JSON(utils.NewCorrectWithGinContext("上传头像成功", ctx).Updated(map[string]any{"avatar": newAvatar}).ToGinResponse())
}

// Detail 详情
func (AccountCtrl) Detail(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		account models.AccountMdl
	)
	ret = models.NewAccountMdl().SetCtx(ctx).GetDbUseQuery("").Where("uuid", ctx.Param("uuid")).First(&account)
	wrongs.ThrowWhenEmpty(ret, "用户")

	ctx.JSON(utils.NewCorrectWithGinContext("", ctx).Datum(map[string]any{"account": account}).ToGinResponse())
}

// List 列表
func (receiver AccountCtrl) List(ctx *gin.Context) {
	var accounts []models.AccountMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForPager(
				models.AccountMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&accounts)
					return map[string]any{"accounts": accounts}
				},
			).ToGinResponse(),
	)
}

// ListJdt jquery-dataTable分页列表
func (receiver AccountCtrl) ListJdt(ctx *gin.Context) {
	var accounts []models.AccountMdl

	ctx.JSON(
		utils.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				models.AccountMdl{}.GetListByQuery(ctx),
				func(db *gorm.DB) map[string]any {
					db.Find(&accounts)
					return map[string]any{"accounts": accounts}
				},
			).ToGinResponse(),
	)
}

// PutUpdatePassword 修改密码
func (recevier AccountCtrl) PutUpdatePassword(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		account *models.AccountMdl
	)

	form := AccountUpdatePasswordForm{}.ShouldBind(ctx)

	ret = models.NewAccountMdl().GetDb("").Where("uuid = ?", ctx.Param("uuid")).First(&account)
	wrongs.ThrowWhenEmpty(ret, "用户")

	// 验证密码
	utils.CheckPassword(form.OldPassword, account.Password)

	account.Password = utils.GeneratePassword(form.Password)
	models.NewAccountMdl().GetDb("").Where("uuid = ?", ctx.Param("uuid")).Save(&account)

	ctx.JSON(utils.NewCorrectWithGinContext("修改成功", ctx).Blank().ToGinResponse())
}
