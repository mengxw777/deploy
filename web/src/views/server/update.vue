<template>
  <div class="app-container">
    <el-form ref="form" :model="form" :rules="rules">
      <el-form-item label="服务器名称" prop="name" label-width="100px">
        <el-input v-model="form.name" placeholder="请输入项目名称"/>
      </el-form-item>
      <el-form-item label="登录IP" prop="ip" label-width="100px">
        <el-input v-model="form.ip" placeholder="请输入项目描述(非必填)"/>
      </el-form-item>
      <el-form-item label="登录端口" prop="port" label-width="100px">
        <el-input v-model="form.port" placeholder="请输入仓库地址"/>
      </el-form-item>
      <el-form-item label="集群" prop="group_id" label-width="100px">
        <el-select v-model="form.group_id" placeholder="请选择集群">
          <el-option
            v-for="item in serverGroupList"
            :key="item.id"
            :label="item.name"
            :value="item.id">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item style="float: right">
        <el-button type="primary" @click="onSubmit('form')">提交</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
  import { getListArr } from '@/api/serverGroup'
  import { show, update } from '@/api/server'

  export default {
    data() {
      return {
        form: {
          name: '',
          ip: '',
          port: '',
          group_id: ''
        },
        serverGroupList: [],
        serverID: '',
        rules: {
          name: [
            { required: true, message: '请输入服务器名称', trigger: 'blur' },
            { min: 1, max: 50, message: '长度在 1 到 50 个字符', trigger: 'blur' }
          ],
          ip: [
            { required: true, message: '请输入服务器IP', trigger: 'blur' }
          ],
          port: [
            { required: true, message: '请输入服务器端口', trigger: 'blur' }
          ]
        }
      }
    },
    created() {
      this.fetchServerGroup()
      this.fetchServer()
    },
    methods: {
      onSubmit(form) {
        this.$refs[form].validate((valid) => {
          if (valid) {

            if (this.form.group_id === '') {
              delete this.form.group_id
            }

            delete this.form.group

            update(this.serverID, this.form).then(response => {
              this.$message({
                message: response.message,
                type: 'success'
              })

              this.$router.push({ name: 'server.index' })
            })
          } else {
            return false
          }
        })
      },
      fetchServerGroup() {
        getListArr().then(response => {
          this.serverGroupList = response.data
        })
      },
      fetchServer() {
        let serverID = this.$route.params.id
        this.serverID = serverID

        show(serverID).then(response => {
          this.form = response.data
        })

      }
    }
  }
</script>

<style scoped>
  .line {
    text-align: center;
  }
</style>

