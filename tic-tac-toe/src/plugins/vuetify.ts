import { createVuetify } from "vuetify";
import * as componets from "vuetify/components"
import * as directives from "vuetify/directives"
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

const vuetify = createVuetify({
  components: componets,
  directives: directives
})

export default vuetify
