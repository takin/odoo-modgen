package template

var CInit = `
# -*- coding: utf-8 -*-

from . import controllers
`

var Controllers = `
# -*- coding: utf-8 -*-
from odoo import http

class {MOD_NAME}(http.Controller):
    @http.route('/{MOD_NAME_LOWER}/{MOD_NAME_LOWER}/', auth='public')
    def index(self, **kw):
        return "Hello, world from {MOD_NAME}"

#     @http.route('/{MOD_NAME}/{MOD_NAME}/objects/', auth='public')
#     def list(self, **kw):
#         return http.request.render('{MOD_NAME_LOWER}.listing', {
#             'root': '/{MOD_NAME_LOWER}/{MOD_NAME_LOWER}',
#             'objects': http.request.env['{MOD_NAME_LOWER}.{MOD_NAME_LOWER}'].search([]),
#         })

#     @http.route('/{MOD_NAME_LOWER}/{MOD_NAME_LOWER}/objects/<model("{MOD_NAME_LOWER}.{MOD_NAME_LOWER}"):obj>/', auth='public')
#     def object(self, obj, **kw):
#         return http.request.render('{MOD_NAME_LOWER}.object', {
#             'object': obj
#         })`
