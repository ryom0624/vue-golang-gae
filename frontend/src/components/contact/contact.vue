<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="6">
        <v-form
          ref="form"
          max-width="400"
        >
          <h1>{{PageTitle}}</h1>
          <v-text-field
            v-model="email"
            :counter="255"
            label="email"
            requird>
          </v-text-field>
          <v-text-field
            v-model="subject"
            :counter="30"
            label="subject"
            requird>
          </v-text-field>
          <v-textarea
            v-model="body"
            rows="8"
            clearable
            clear-icon="cancel"
            label="body"
          >
          </v-textarea>
          <v-btn @click="post()">Submit</v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from 'axios'

export default {
  name: 'contact',
  data () {
    return {
      PageTitle: 'CONTACT',
      email: '',
      subject: '',
      body: ''
    }
  },
  methods: {
    post () {
      let contactData
      contactData = {
        email: this.email,
        subject: this.subject,
        body: this.body
      }
      console.log('contactData', contactData)
      axios.post('http://localhost:8080/api/v1/contact', contactData).then(res => {
        console.log('contactDdata', contactData)
        console.log('%cPOST REQUEST data :%c' + contactData, 'color:red; font-weight:bold;', '')
        console.log('%cPOST contact RESPONSE data :%c' + JSON.stringify(res.data, null, 4), 'color:red; font-weight:bold;', '')
      })
      this.$router.push(`/thanks`)
    }
  }
}

</script>
