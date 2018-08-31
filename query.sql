-- 用户
insert into sdb_user(account,pass,phone) value("admin","0192023a7bbd73250516f069df18b500","0");
insert into sdb_user(account,pass,phone) value("admin1","0192023a7bbd73250516f069df18b500","1")

-- 权限基础数据
insert into sdb_permissions(`name`,img,created_id,is_show,pid,url,pn)VALUES
("权限管理","fa-id-card",1,1,0,"",""),
("用户管理","fa-user-o",1,1,1,"/user/list","user"),
("角色管理","fa-user-circle-o",1,1,1,"/role/list","role")
-- 角色
insert into sdb_role(permission_ids,`name`,detail) VALUE("1,2,3","管理员","管理员")