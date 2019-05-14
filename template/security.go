package template


var Security = `
id,name,model_id:id,group_id:id,perm_read,perm_write,perm_create,perm_unlink
access_{MOD_NAME_LOWER}_{MOD_NAME_LOWER},{MOD_NAME_LOWER}.{MOD_NAME_LOWER},model_{MOD_NAME_LOWER}_{MOD_NAME_LOWER},,1,0,0,0
`