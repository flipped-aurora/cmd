//+build postgres

package data

import (
	"github.com/flipped-aurora/gva/library/global"
	model "github.com/flipped-aurora/gva/model/gfva/system"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var Casbin = new(casbin)

type casbin struct{}

// Init casbin_rule 表数据初始化
// Author: [SliverHorn](https://github.com/SliverHorn)
func (c *casbin) Init() error {
	carbines := []model.Casbin{
		{PType: "p", AuthorityId: "888", Path: "/base/login", Method: "POST"},

		{PType: "p", AuthorityId: "888", Path: "/user/register", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/user/changePassword", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/user/setUserAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/user/setUserInfo", Method: "PUT"},
		{PType: "p", AuthorityId: "888", Path: "/user/deleteUser", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/user/getUserList", Method: "POST"},

		{PType: "p", AuthorityId: "888", Path: "/api/createApi", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/api/getApiById", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/api/updateApi", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/api/deleteApi", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/api/getApiList", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/api/getAllApis", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/api/deleteApisByIds", Method: "DELETE"},

		{PType: "p", AuthorityId: "888", Path: "/authority/createAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/authority/copyAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/authority/updateAuthority", Method: "PUT"},
		{PType: "p", AuthorityId: "888", Path: "/authority/setDataAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/authority/deleteAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/authority/getAuthorityList", Method: "POST"},

		{PType: "p", AuthorityId: "888", Path: "/menu/addBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/getBaseMenuById", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/updateBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/deleteBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/getBaseMenuTree", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/addMenuAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/getMenu", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/getMenuList", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/getMenuAuthority", Method: "POST"},

		{PType: "p", AuthorityId: "888", Path: "/fileUploadAndDownload/upload", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/fileUploadAndDownload/deleteFile", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/fileUploadAndDownload/getFileList", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/fileUploadAndDownload/findFile", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/fileUploadAndDownload/removeChunk", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/fileUploadAndDownload/breakpointContinue", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/fileUploadAndDownload/breakpointContinueFinish", Method: "POST"},

		{PType: "p", AuthorityId: "888", Path: "/casbin/updateCasbin", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/casbin/getPolicyPathByAuthorityId", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/casbin/casbinTest/:pathParam", Method: "GET"},

		{PType: "p", AuthorityId: "888", Path: "/jwt/jsonInBlacklist", Method: "POST"},

		{PType: "p", AuthorityId: "888", Path: "/system/getSystemConfig", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/system/setSystemConfig", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/system/getServerInfo", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/system/reloadSystem", Method: "POST"},

		{PType: "p", AuthorityId: "888", Path: "/customer/customer", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/customer/customer", Method: "PUT"},
		{PType: "p", AuthorityId: "888", Path: "/customer/customer", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/customer/customer", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/customer/customerList", Method: "GET"},

		{PType: "p", AuthorityId: "888", Path: "/autoCode/createTemp", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/autoCode/getTables", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/autoCode/getDB", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/autoCode/getColumn", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/autoCode/preview", Method: "POST"},

		{PType: "p", AuthorityId: "888", Path: "/sysDictionary/createSysDictionary", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionary/findSysDictionary", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionary/updateSysDictionary", Method: "PUT"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionary/deleteSysDictionary", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionary/getSysDictionaryList", Method: "GET"},

		{PType: "p", AuthorityId: "888", Path: "/sysDictionaryDetail/createSysDictionaryDetail", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionaryDetail/findSysDictionaryDetail", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Method: "PUT"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Method: "GET"},

		{PType: "p", AuthorityId: "888", Path: "/sysOperationRecord/createSysOperationRecord", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/sysOperationRecord/findSysOperationRecord", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/sysOperationRecord/deleteSysOperationRecord", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/sysOperationRecord/updateSysOperationRecord", Method: "PUT"},
		{PType: "p", AuthorityId: "888", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/sysOperationRecord/getSysOperationRecordList", Method: "GET"},

		{PType: "p", AuthorityId: "888", Path: "/email/emailTest", Method: "POST"},

		{PType: "p", AuthorityId: "888", Path: "/simpleUploader/upload", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/simpleUploader/checkFileMd5", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/simpleUploader/mergeFileMd5", Method: "GET"},

		{PType: "p", AuthorityId: "888", Path: "/workflowProcess/createWorkflowProcess", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/workflowProcess/deleteWorkflowProcess", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/workflowProcess/deleteWorkflowProcessByIds", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/workflowProcess/updateWorkflowProcess", Method: "PUT"},
		{PType: "p", AuthorityId: "888", Path: "/workflowProcess/findWorkflowProcess", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/workflowProcess/getWorkflowProcessList", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/workflowProcess/findWorkflowStep", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/workflowProcess/startWorkflow", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/workflowProcess/completeWorkflowMove", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/workflowProcess/getMyStated", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/workflowProcess/getMyNeed", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/workflowProcess/getWorkflowMoveByID", Method: "GET"},

		{PType: "p", AuthorityId: "8881", Path: "/base/login", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/user/register", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/api/createApi", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/api/getApiList", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/api/getApiById", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/api/deleteApi", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/api/updateApi", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/api/getAllApis", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/authority/createAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/authority/deleteAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/authority/getAuthorityList", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/authority/setDataAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/getMenu", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/getMenuList", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/addBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/getBaseMenuTree", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/addMenuAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/getMenuAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/deleteBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/updateBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/getBaseMenuById", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/user/changePassword", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/user/getUserList", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/user/setUserAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/fileUploadAndDownload/upload", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/fileUploadAndDownload/getFileList", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/fileUploadAndDownload/deleteFile", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/casbin/updateCasbin", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/casbin/getPolicyPathByAuthorityId", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/jwt/jsonInBlacklist", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/system/getSystemConfig", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/system/setSystemConfig", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/customer/customer", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/customer/customer", Method: "PUT"},
		{PType: "p", AuthorityId: "8881", Path: "/customer/customer", Method: "DELETE"},
		{PType: "p", AuthorityId: "8881", Path: "/customer/customer", Method: "GET"},
		{PType: "p", AuthorityId: "8881", Path: "/customer/customerList", Method: "GET"},

		{PType: "p", AuthorityId: "9528", Path: "/base/login", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/user/register", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/api/createApi", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/api/getApiList", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/api/getApiById", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/api/deleteApi", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/api/updateApi", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/api/getAllApis", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/authority/createAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/authority/deleteAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/authority/getAuthorityList", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/authority/setDataAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/getMenu", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/getMenuList", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/addBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/getBaseMenuTree", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/addMenuAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/getMenuAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/deleteBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/updateBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/getBaseMenuById", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/user/changePassword", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/user/getUserList", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/user/setUserAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/fileUploadAndDownload/upload", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/fileUploadAndDownload/getFileList", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/fileUploadAndDownload/deleteFile", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/casbin/updateCasbin", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/casbin/getPolicyPathByAuthorityId", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/jwt/jsonInBlacklist", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/system/getSystemConfig", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/system/setSystemConfig", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/customer/customer", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/customer/customer", Method: "PUT"},
		{PType: "p", AuthorityId: "9528", Path: "/customer/customer", Method: "DELETE"},
		{PType: "p", AuthorityId: "9528", Path: "/customer/customer", Method: "GET"},
		{PType: "p", AuthorityId: "9528", Path: "/customer/customerList", Method: "GET"},
		{PType: "p", AuthorityId: "9528", Path: "/autoCode/createTemp", Method: "POST"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Find(&[]model.Casbin{}).RowsAffected >= 158 {
			color.Danger.Println("\n[Mysql] --> casbin_rule 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&carbines).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

// TableName 定义表名
// Author: [SliverHorn](https://github.com/SliverHorn)
func (c *casbin) TableName() string {
	return "casbin_rule"
}
