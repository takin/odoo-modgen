package template

var ViewTemplates = `
<odoo>
    <data>
        <!-- <template id="listing"> -->
        <!--   <ul> -->
        <!--     <li t-foreach="objects" t-as="object"> -->
        <!--       <a t-attf-href="#{ root }/objects/#{ object.id }"> -->
        <!--         <t t-esc="object.display_name"/> -->
        <!--       </a> -->
        <!--     </li> -->
        <!--   </ul> -->
        <!-- </template> -->
        <!-- <template id="object"> -->
        <!--   <h1><t t-esc="object.display_name"/></h1> -->
        <!--   <dl> -->
        <!--     <t t-foreach="object._fields" t-as="field"> -->
        <!--       <dt><t t-esc="field"/></dt> -->
        <!--       <dd><t t-esc="object[field]"/></dd> -->
        <!--     </t> -->
        <!--   </dl> -->
        <!-- </template> -->
    </data>
</odoo>
`

var ViewViews = `
<odoo>
  <data>
    <!-- explicit list view definition -->
    <!--
    <record model="ir.ui.view" id="{MOD_NAME_LOWER}.list">
      <field name="name">{MOD_NAME_LOWER} list</field>
      <field name="model">{MOD_NAME_LOWER}.{MOD_NAME_LOWER}</field>
      <field name="arch" type="xml">
        <tree>
          <field name="name"/>
          <field name="value"/>
          <field name="value2"/>
        </tree>
      </field>
    </record>
    -->

    <!-- actions opening views on models -->
    <!--
    <record model="ir.actions.act_window" id="{MOD_NAME_LOWER}.action_window">
      <field name="name">{MOD_NAME_LOWER} window</field>
      <field name="res_model">{MOD_NAME_LOWER}.{MOD_NAME_LOWER}</field>
      <field name="view_mode">tree,form</field>
    </record>
    -->

    <!-- server action to the one above -->
    <!--
    <record model="ir.actions.server" id="{MOD_NAME_LOWER}.action_server">
      <field name="name">{MOD_NAME_LOWER} server</field>
      <field name="model_id" ref="model_{MOD_NAME_LOWER}_{MOD_NAME_LOWER}"/>
      <field name="state">code</field>
      <field name="code">
        action = {
          "type": "ir.actions.act_window",
          "view_mode": "tree,form",
          "res_model": self._name,
        }
      </field>
    </record>
    -->

    <!-- Top menu item -->
    <!--
    <menuitem name="{MOD_NAME_LOWER}" id="{MOD_NAME_LOWER}.menu_root"/>
    -->
    <!-- menu categories -->
    <!--
    <menuitem name="Menu 1" id="{MOD_NAME_LOWER}.menu_1" parent="{MOD_NAME_LOWER}.menu_root"/>
    <menuitem name="Menu 2" id="{MOD_NAME_LOWER}.menu_2" parent="{MOD_NAME_LOWER}.menu_root"/>
    -->
    <!-- actions -->
    <!--
    <menuitem name="List" id="{MOD_NAME_LOWER}.menu_1_list" parent="{MOD_NAME_LOWER}.menu_1"
              action="{MOD_NAME_LOWER}.action_window"/>
    <menuitem name="Server to list" id="{MOD_NAME_LOWER}" parent="{MOD_NAME_LOWER}.menu_2"
              action="{MOD_NAME_LOWER}.action_server"/>
    -->
  </data>
</odoo>
`
