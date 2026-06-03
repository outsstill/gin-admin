# Gin Admin Core API 文档

本文档根据 `routes/admin.go`、`routes/static.go`、对应控制器和请求结构体整理生成。

生成日期：2026-05-26

## 基础信息

- 默认接口前缀：`/admin`
- 前缀来源：`admin.Register(r, prefix, modules...)`，当 `prefix == ""` 时使用 `/admin`
- 默认请求格式：`application/json`
- 文件上传格式：`multipart/form-data`
- 认证请求头：`Authorization: Bearer <token>`

## 鉴权说明

`/admin` 路由组默认使用：

- `AuthAdminJWT()`
- `OperationLog()`

以下路径在当前代码中跳过 JWT 鉴权：

| 路径 | 说明 |
| --- | --- |
| `/admin/auth/login` | 登录 |
| `/admin/auth/captcha` | 验证码 |
| `/admin/upload` | 文件上传 |
| `/admin/version` | 版本号 |
| `/admin/test` | 预留路径，当前路由文件未注册 |

以下路径需要 JWT，但跳过权限表校验：

| 路径 | 说明 |
| --- | --- |
| `/admin/auth/logout` | 退出登录 |
| `/admin/auth/refresh-token` | 刷新令牌 |
| `/admin/auth/current` | 当前用户 |
| `/admin/roles/all` | 全量角色 |
| `/admin/permissions/all` | 全量权限 |
| `/admin/menus/all` | 全量菜单 |

除上述路径外，请求需要有效 Admin Token，并且当前用户必须是超级管理员，或拥有匹配当前 `method + path` 的权限记录。

## 通用响应

成功响应：

```json
{
  "code": 200,
  "success": true,
  "message": "操作成功",
  "data": {}
}
```

无数据成功响应：

```json
{
  "code": 200,
  "success": true,
  "message": "操作成功"
}
```

业务失败响应通常为 HTTP 200：

```json
{
  "code": 400,
  "success": false,
  "message": "请求处理失败"
}
```

参数校验失败响应通常为 HTTP 200：

```json
{
  "code": 422,
  "message": "请求验证不通过，具体请查看 errors",
  "errors": {
    "FieldName": ["error message"]
  }
}
```

Token 失效响应通常为 HTTP 200，业务码为 `701`。

## 通用分页参数

分页接口返回：

```json
{
  "data": [],
  "pager": {
    "CurrentPage": 1,
    "PerPage": 10,
    "TotalPage": 1,
    "TotalCount": 1,
    "NextPageURL": "",
    "PrevPageURL": ""
  }
}
```

默认分页参数名来自配置文件，`example.config.yaml` 中为：

| 参数 | 类型 | 默认值 | 说明 |
| --- | --- | --- | --- |
| `page` | int | `1` | 页码 |
| `per_page` | int | `10` | 每页数量 |
| `sort` | string | `id` | 排序字段 |
| `order` | string | `asc` | 排序方向 |

## 路由总览

| 方法 | 路径 | 控制器方法 | 鉴权 |
| --- | --- | --- | --- |
| GET | `/admin/index` | `AdminIndexController.Index` | JWT + 权限 |
| GET | `/admin/version` | `AdminIndexController.Version` | 无需 JWT |
| GET | `/admin/setting-all` | `AdminConfigController.AllShow` | JWT + 权限 |
| POST | `/admin/auth/login` | `AdminAuthController.Login` | 游客 |
| POST | `/admin/auth/refresh-token` | `AdminAuthController.RefreshToken` | JWT |
| GET | `/admin/auth/current` | `AdminAuthController.Current` | JWT |
| POST | `/admin/auth/profile` | `AdminAuthController.UpdateProfile` | JWT + 权限 |
| POST | `/admin/auth/profile-pass` | `AdminAuthController.UpdatePassword` | JWT + 权限 |
| POST | `/admin/auth/logout` | `AdminAuthController.Logout` | JWT |
| GET | `/admin/auth/captcha` | `AdminAuthController.ShowCaptcha` | 无需 JWT |
| GET | `/admin/users` | `AdminUserController.Index` | JWT + 权限 |
| GET | `/admin/user/:id` | `AdminUserController.Get` | JWT + 权限 |
| POST | `/admin/user` | `AdminUserController.Store` | JWT + 权限 |
| PUT | `/admin/user/:id` | `AdminUserController.Update` | JWT + 权限 |
| DELETE | `/admin/user/:id` | `AdminUserController.Delete` | JWT + 权限 |
| GET | `/admin/roles` | `AdminRoleController.Index` | JWT + 权限 |
| GET | `/admin/roles/all` | `AdminRoleController.All` | JWT |
| GET | `/admin/role/:id` | `AdminRoleController.Get` | JWT + 权限 |
| POST | `/admin/role` | `AdminRoleController.Store` | JWT + 权限 |
| PUT | `/admin/role/:id` | `AdminRoleController.Update` | JWT + 权限 |
| DELETE | `/admin/role/:id` | `AdminRoleController.Delete` | JWT + 权限 |
| GET | `/admin/menus` | `AdminMenuController.Index` | JWT + 权限 |
| GET | `/admin/menus/all` | `AdminMenuController.All` | JWT |
| GET | `/admin/menu/:id` | `AdminMenuController.Get` | JWT + 权限 |
| POST | `/admin/menu` | `AdminMenuController.Store` | JWT + 权限 |
| PUT | `/admin/menu/:id` | `AdminMenuController.Update` | JWT + 权限 |
| DELETE | `/admin/menu/:id` | `AdminMenuController.Delete` | JWT + 权限 |
| GET | `/admin/permissions` | `AdminPermissionController.Index` | JWT + 权限 |
| GET | `/admin/permissions/all` | `AdminPermissionController.All` | JWT |
| GET | `/admin/permission/:id` | `AdminPermissionController.Get` | JWT + 权限 |
| POST | `/admin/permission` | `AdminPermissionController.Store` | JWT + 权限 |
| PUT | `/admin/permission/:id` | `AdminPermissionController.Update` | JWT + 权限 |
| DELETE | `/admin/permission/:id` | `AdminPermissionController.Delete` | JWT + 权限 |
| GET | `/admin/configs` | `AdminConfigController.Index` | JWT + 权限 |
| GET | `/admin/configs/all` | `AdminConfigController.All` | JWT + 权限 |
| GET | `/admin/config/:id` | `AdminConfigController.Get` | JWT + 权限 |
| POST | `/admin/config` | `AdminConfigController.Store` | JWT + 权限 |
| PUT | `/admin/config/:id` | `AdminConfigController.Update` | JWT + 权限 |
| DELETE | `/admin/config/:id` | `AdminConfigController.Delete` | JWT + 权限 |
| POST | `/admin/upload` | `AdminFileController.Upload` | 无需 JWT |
| POST | `/admin/file` | `AdminFileController.Store` | JWT + 权限 |
| GET | `/admin/files` | `AdminFileController.Index` | JWT + 权限 |
| GET | `/admin/file/:id` | `AdminFileController.Get` | JWT + 权限 |
| PUT | `/admin/file/:id` | `AdminFileController.Update` | JWT + 权限 |
| DELETE | `/admin/file/:id` | `AdminFileController.Delete` | JWT + 权限 |
| POST | `/admin/log` | `AdminOperationController.Store` | JWT + 权限 |
| GET | `/admin/logs` | `AdminOperationController.Index` | JWT + 权限 |
| GET | `/admin/log/:id` | `AdminOperationController.Get` | JWT + 权限 |
| PUT | `/admin/log/:id` | `AdminOperationController.Update` | JWT + 权限 |
| DELETE | `/admin/log/:id` | `AdminOperationController.Delete` | JWT + 权限 |

## 系统接口

### GET `/admin/index`

返回后台首页检查结果。

响应：

```json
{
  "code": 200,
  "msg": "ok"
}
```

### GET `/admin/version`

返回配置中的应用版本号。

响应：

```json
{
  "code": 200,
  "data": "1.0.0"
}
```

### GET `/admin/setting-all`

返回允许前端读取的配置项，即 `configs.is_can_front = 1` 的记录。

## 认证接口

### POST `/admin/auth/login`

登录并返回 Admin Token。

请求体：

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `username` | string | 是 | 用户名 |
| `password` | string | 是 | 密码 |
| `captcha_id` | string | 否 | 验证码 ID；结构体中声明了字段，但当前登录方法未启用验证码校验 |
| `captcha_answer` | string | 否 | 验证码答案；结构体中声明了字段，但当前登录方法未启用验证码校验 |

响应：

```json
{
  "code": 200,
  "success": true,
  "message": "操作成功",
  "data": {
    "token": "<jwt>"
  }
}
```

### POST `/admin/auth/refresh-token`

刷新当前 Token。Token 剩余有效期大于 5 分钟时，当前实现会返回刷新失败。

响应：

```json
{
  "data": {
    "token": "<new-jwt>"
  }
}
```

### GET `/admin/auth/current`

返回当前登录管理员。

查询参数：

| 参数 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `menus_on` | int | 否 | 大于 0 时附带用户菜单和 `dashboardMenus` |

### POST `/admin/auth/profile`

更新当前管理员资料。

请求体：

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `name` | string | 是 | 昵称 |

### POST `/admin/auth/profile-pass`

更新当前管理员密码。

请求体：

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `password` | string | 是 | 新密码，最少 6 位 |
| `confirm_password` | string | 是 | 确认密码，必须与 `password` 一致 |

### POST `/admin/auth/logout`

退出登录。

当前实现直接返回成功，不会服务端注销 Token。

### GET `/admin/auth/captcha`

生成图形验证码。

响应：

```json
{
  "data": {
    "captcha_id": "<id>",
    "captcha_image": "<base64-image>"
  }
}
```

## 管理员账号

账号模型主要字段：`id`、`username`、`name`、`avatar_id`、`avatar`、`roles`、`created_at`、`updated_at`。

### GET `/admin/users`

分页获取管理员账号。

查询参数：通用分页参数。

### GET `/admin/user/:id`

获取单个管理员账号。

路径参数：

| 参数 | 类型 | 说明 |
| --- | --- | --- |
| `id` | uint64 | 管理员账号 ID |

### POST `/admin/user`

创建管理员账号。

请求体：

| 字段 | 类型 | 必填 | 校验/说明 |
| --- | --- | --- | --- |
| `username` | string | 是 | 唯一 |
| `password` | string | 是 | 最少 6 位 |
| `confirm_password` | string | 是 | 必须与 `password` 一致 |
| `name` | string | 是 | 昵称 |
| `role_ids` | uint[] | 否 | 关联角色 ID |
| `avatar_id` | uint | 否 | 头像文件 ID |

### PUT `/admin/user/:id`

更新管理员账号。

请求体：

| 字段 | 类型 | 必填 | 校验/说明 |
| --- | --- | --- | --- |
| `username` | string | 否 | 唯一，排除当前账号 |
| `password` | string | 否 | 最少 6 位 |
| `name` | string | 否 | 昵称 |
| `role_ids` | uint[] | 否 | 替换账号角色 |
| `avatar_id` | uint | 否 | 头像文件 ID |

### DELETE `/admin/user/:id`

删除管理员账号。

## 角色

角色模型主要字段：`id`、`name`、`slug`、`permissions`、`menus`、`created_at`、`updated_at`。

### GET `/admin/roles`

分页获取角色。

查询参数：通用分页参数。

### GET `/admin/roles/all`

获取全部角色。

### GET `/admin/role/:id`

获取单个角色，并预加载 `menus` 和 `permissions`。

### POST `/admin/role`

创建角色。

请求体：

| 字段 | 类型 | 必填 | 校验/说明 |
| --- | --- | --- | --- |
| `name` | string | 是 | 角色名称 |
| `slug` | string | 是 | 唯一标识 |
| `permission_ids` | uint[] | 否 | 关联权限 ID |
| `menu_ids` | uint[] | 否 | 关联菜单 ID |

### PUT `/admin/role/:id`

更新角色。

请求体：

| 字段 | 类型 | 必填 | 校验/说明 |
| --- | --- | --- | --- |
| `name` | string | 是 | 角色名称 |
| `slug` | string | 是 | 唯一标识，排除当前角色 |
| `permission_ids` | uint[] | 否 | 替换关联权限 |
| `menu_ids` | uint[] | 否 | 替换关联菜单 |

### DELETE `/admin/role/:id`

删除角色。

## 菜单

菜单模型主要字段：`id`、`parent_id`、`order`、`name`、`icon`、`path`、`uri`、`created_at`、`updated_at`。

### GET `/admin/menus`

分页获取菜单。

查询参数：通用分页参数。

### GET `/admin/menus/all`

获取全部菜单。

### GET `/admin/menu/:id`

获取单个菜单。

### POST `/admin/menu`

创建菜单。

请求体：

| 字段 | 类型 | 必填 | 校验/说明 |
| --- | --- | --- | --- |
| `name` | string | 是 | 菜单名称，唯一 |
| `order` | int64 | 是 | 排序值 |
| `path` | string | 是 | 访问路径 |
| `uri` | string | 是 | 前端组件或 URI |
| `parent_id` | uint64 | 否 | 父菜单 ID |
| `icon` | string | 否 | 图标 |

### PUT `/admin/menu/:id`

更新菜单。

请求体：

| 字段 | 类型 | 必填 | 校验/说明 |
| --- | --- | --- | --- |
| `name` | string | 否 | 菜单名称，唯一，排除当前菜单 |
| `order` | int64 | 否 | 排序值 |
| `parent_id` | uint64 | 否 | 父菜单 ID |
| `icon` | string | 否 | 图标 |
| `uri` | string | 否 | 前端组件或 URI |
| `path` | string | 否 | 访问路径 |

### DELETE `/admin/menu/:id`

删除菜单。

## 权限

权限模型主要字段：`id`、`name`、`slug`、`http_method`、`http_path`、`order`、`parent_id`、`created_at`、`updated_at`。

### GET `/admin/permissions`

分页获取权限。

查询参数：通用分页参数。

### GET `/admin/permissions/all`

获取全部权限。

### GET `/admin/permission/:id`

获取单个权限。

### POST `/admin/permission`

创建权限。

请求体：

| 字段 | 类型 | 必填 | 校验/说明 |
| --- | --- | --- | --- |
| `name` | string | 是 | 权限名称 |
| `slug` | string | 是 | 唯一标识 |
| `http_method` | string | 是 | HTTP 方法；保存时转为小写，`any` 表示不限方法 |
| `http_path` | string | 是 | 匹配路由路径 |
| `order` | uint64 | 是 | 排序值 |
| `parent_id` | uint64 | 否 | 父权限 ID |

### PUT `/admin/permission/:id`

更新权限。

请求体：

| 字段 | 类型 | 必填 | 校验/说明 |
| --- | --- | --- | --- |
| `name` | string | 是 | 权限名称 |
| `slug` | string | 是 | 唯一标识，排除当前权限 |
| `http_method` | string | 否 | HTTP 方法；保存时转为小写 |
| `http_path` | string | 否 | 匹配路由路径 |
| `order` | uint64 | 是 | 排序值 |
| `parent_id` | uint64 | 否 | 父权限 ID |

### DELETE `/admin/permission/:id`

删除权限。

## 配置

配置模型主要字段：`id`、`config_key`、`config_value`、`config_label`、`type`、`options`、`describe`、`is_can_front`、`is_required`、`order`、`group_id`、`state`、`show_type`、`placeholder`、`created_at`、`updated_at`。

### GET `/admin/configs`

分页获取配置。

查询参数：

| 参数 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `page` | int | 否 | 页码 |
| `per_page` | int | 否 | 每页数量 |
| `sort` | string | 否 | 排序字段 |
| `order` | string | 否 | 排序方向 |
| `config_key` | string | 否 | 按 `config_key` 前缀模糊查询 |
| `config_label` | string | 否 | 按 `config_label` 前缀模糊查询 |

### GET `/admin/configs/all`

获取全部配置。

### GET `/admin/config/:id`

获取单个配置。

### POST `/admin/config`

创建配置。

请求体：

| 字段 | 类型 | 必填 | 校验/说明 |
| --- | --- | --- | --- |
| `config_key` | string | 是 | 配置键，唯一 |
| `config_value` | string | 是 | 配置值 |
| `config_label` | string | 是 | 配置标签 |
| `type` | int | 否 | 类型 |
| `options` | string | 否 | 选项 |
| `describe` | string | 否 | 描述 |
| `is_can_front` | int | 否 | 是否允许前端读取 |
| `is_required` | uint | 否 | 是否必填 |
| `order` | uint | 否 | 排序值 |
| `group_id` | uint | 否 | 分组 ID |
| `state` | uint | 否 | 状态 |
| `show_type` | string | 否 | 展示类型 |
| `placeholder` | string | 否 | 占位提示 |

### PUT `/admin/config/:id`

更新配置。

请求体：

| 字段 | 类型 | 必填 | 校验/说明 |
| --- | --- | --- | --- |
| `config_key` | string | 否 | 配置键，唯一，排除当前配置 |
| `config_value` | string | 是 | 配置值 |
| `config_label` | string | 否 | 配置标签 |
| `type` | int | 否 | 类型 |
| `options` | string | 否 | 选项 |
| `describe` | string | 否 | 描述 |
| `is_can_front` | int | 否 | 是否允许前端读取 |
| `is_required` | uint | 否 | 是否必填 |
| `order` | uint | 否 | 排序值 |
| `group_id` | uint | 否 | 分组 ID |
| `state` | uint | 否 | 状态 |
| `show_type` | string | 否 | 展示类型 |
| `placeholder` | string | 否 | 占位提示 |

### DELETE `/admin/config/:id`

删除配置。

## 文件

文件模型主要响应字段：`id`、`origin_name`、`name`、`key`、`group_id`、`size`、`storage`、`type`、`ext`、`url`、`content_type`、`e_tag`、`bucket`、`last_modified`、`full_url`、`deleted_at`、`created_at`、`updated_at`。

### POST `/admin/upload`

上传文件并创建文件记录。

请求格式：`multipart/form-data`

表单字段：

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `file` | file | 是 | 上传文件 |
| `uploadStorage` | string | 否 | 指定存储驱动；为空时使用配置 `storage.driver` |
| `group_id` | int | 否 | 文件分组 ID，默认 `99` |
| `type` | int | 否 | 文件类型，默认 `1` |

限制来自配置：

- 文件大小：`storage.size_limit`
- 文件后缀：`storage.ext`

### POST `/admin/file`

手动创建文件记录。

请求体：

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `origin_name` | string | 否 | 原始文件名 |
| `name` | string | 否 | 文件名 |
| `size` | int64 | 否 | 文件大小 |
| `storage` | string | 否 | 存储驱动 |
| `path` | string | 否 | 存储路径 |
| `type` | int | 否 | 文件类型 |
| `ext` | string | 否 | 后缀 |
| `url` | string | 否 | 文件 URL |

当前创建方法会将 `last_modified` 设置为当前时间，并将 `user_id` 设置为当前管理员 ID。

### GET `/admin/files`

分页获取文件。

查询参数：

| 参数 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `page` | int | 否 | 页码 |
| `per_page` | int | 否 | 每页数量 |
| `sort` | string | 否 | 排序字段 |
| `order` | string | 否 | 排序方向 |
| `storage` | string | 否 | 按存储驱动精确查询 |

### GET `/admin/file/:id`

获取单个文件。

### PUT `/admin/file/:id`

更新文件记录。

请求体：

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `origin_name` | string | 否 | 原始文件名 |
| `name` | string | 否 | 文件名 |
| `key` | string | 否 | 对象 Key |
| `group_id` | int | 否 | 分组 ID |
| `size` | int64 | 否 | 文件大小 |
| `storage` | string | 否 | 存储驱动 |
| `path` | string | 否 | 存储路径 |
| `type` | int | 否 | 文件类型 |
| `ext` | string | 否 | 后缀 |
| `user_id` | uint64 | 否 | 用户 ID |
| `url` | string | 否 | 文件 URL |
| `content_type` | string | 否 | MIME 类型 |
| `e_tag` | string | 否 | ETag |
| `bucket` | string | 否 | Bucket |
| `last_modified` | time | 否 | 最后修改时间 |
| `deleted_at` | time | 否 | 删除时间 |

### DELETE `/admin/file/:id`

删除文件。当前实现会先按文件记录中的存储驱动删除实际文件，再删除数据库记录。

## 操作日志

操作日志模型主要字段：`id`、`admin_user`、`user_id`、`path`、`url`、`method`、`ip`、`input`、`created_at`、`updated_at`。

### POST `/admin/log`

创建操作日志。当前实现只写入当前管理员 `user_id`。

### GET `/admin/logs`

分页获取操作日志。

查询参数：

| 参数 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `page` | int | 否 | 页码 |
| `per_page` | int | 否 | 每页数量 |
| `sort` | string | 否 | 排序字段 |
| `order` | string | 否 | 排序方向 |
| `path` | string | 否 | 按请求路径前缀模糊查询 |
| `ip` | string | 否 | 按 IP 前缀模糊查询 |

### GET `/admin/log/:id`

获取单条操作日志，并预加载 `admin_user`。

### PUT `/admin/log/:id`

当前路由指向 `AdminOperationController.Update`，但实现中读取和更新的是 `config` 模型，并使用 `ConfigModelUpdateRequest` 请求体。请确认该行为是否符合预期。

### DELETE `/admin/log/:id`

当前路由指向 `AdminOperationController.Delete`，但实现中删除的是 `config` 模型。请确认该行为是否符合预期。

## 静态资源路由

`routes/static.go` 计划根据配置注册静态文件路由：

- 默认静态前缀：`/static`
- 配置来源：`setting.GlobalSetting.Storage.Local.StaticPrefix`
- 本地目录：`storage/files`

按 Gin 的 `Static` 规则，实际访问形式通常为：

```text
GET /static/*filepath
```

实现备注：当前 `RegisterStaticRoutes` 中只有 `HasRoute(r, staticRoute)` 为 `true` 时才调用 `r.Static(...)`，因此默认情况下可能不会注册新的静态资源路由。请结合运行时路由确认实际行为。
