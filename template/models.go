package template

var MInit = `
# -*- coding: utf-8 -*-

from . import models
`

var Model = `
# -*- coding: utf-8 -*-

from odoo import models, fields, api

class {MOD_NAME}(models.Model):
  _name = '{MOD_NAME_LOWER}.{MOD_NAME_LOWER}'
  
  name = fields.Char()
  value = fields.Integer()
  value2 = fields.Float(compute="_value_pc", store=True)
  description = fields.Text()
  #
  # @api.depends('value')
  # def _value_pc(self):
	#   self.value2 = float(self.value) / 100
`
