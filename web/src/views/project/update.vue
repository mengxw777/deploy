<template>
  <div class="app-container">
    <el-form ref="form" :model="form" :rules="rules">
      <el-form-item label="项目名称" prop="name" label-width="90px">
        <el-input v-model="form.name" placeholder="请输入项目名称"/>
      </el-form-item>
      <el-form-item label="项目描述" label-width="90px">
        <el-input v-model="form.description" placeholder="请输入项目描述(非必填)"/>
      </el-form-item>
      <el-form-item label="仓库地址" prop="git_addr" label-width="90px">
        <el-input v-model="form.git_addr" placeholder="请输入仓库地址"/>
      </el-form-item>
      <el-form-item label="分支名称" prop="branch" label-width="90px">
        <el-input v-model="form.branch" placeholder="请输入分支名称"/>
      </el-form-item>
      <el-form-item label="部署模式" prop="deploy_mode" label-width="90px">
        <el-select v-model="form.deploy_mode" placeholder="请选择部署模式">
          <el-option label="分支" :value="1"/>
          <el-option label="标签" :value="2"/>
        </el-select>
      </el-form-item>
      <el-form-item label="集群" prop="group_id" label-width="90px">
        <el-select v-model="form.group_id" filterable placeholder="请选择集群">
          <el-option
            v-for="item in serverGroupList"
            :key="item.id"
            :label="item.name"
            :value="item.id">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="部署用户" prop="deploy_user" label-width="90px">
        <el-input v-model="form.deploy_user" placeholder="可免密登录部署服务器的用户"/>
      </el-form-item>
      <el-form-item label="部署路径" prop="deploy_path" label-width="90px">
        <el-input v-model="form.deploy_path" placeholder="部署服务器目录"/>
      </el-form-item>
      <el-form-item label="部署前命令" label-width="90px">
        <el-input v-model="form.deploy_before_cmd" type="textarea" placeholder="此命令在部署服务器上执行，一行一条"/>
      </el-form-item>
      <el-form-item label="部署后命令" label-width="90px">
        <el-input v-model="form.deploy_after_cmd" type="textarea" placeholder="此命令在部署服务器上执行，一行一条"/>
      </el-form-item>
      <el-form-item label="构建命令" label-width="90px">
        <el-input v-model="form.build_cmd" type="textarea" placeholder="此命令在构建服务器上执行，一行一条"/>
      </el-form-item>
      <el-form-item style="float: right">
        <template slot-scope="scope">
          <el-button type="primary" @click="onSubmit('form')">提交</el-button>
        </template>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
  import { getList } from '@/api/serverGroup'
  import { show, update } from '@/api/project'

  export default {
    data() {
      return {
        form: {
          name: '',
          description: '',
          git_addr: '',
          branch: '',
          deploy_mode: '',
          group_id: '',
          deploy_user: '',
          deploy_path: '',
          deploy_before_cmd: '',
          deploy_after_cmd: '',
          build_cmd: ''
        },
        serverGroupList: [],
        rules: {
          name: [
            { required: true, message: '请输入活动名称', trigger: 'blur' },
            { min: 1, max: 50, message: '长度在 1 到 50 个字符', trigger: 'blur' }
          ],
          git_addr: [
            { required: true, message: '请输入仓库地址', trigger: 'blur' }
          ],
          branch: [
            { required: true, message: '请输入活动分支名称', trigger: 'blur' }
          ],
          deploy_mode: [
            { required: true, message: '请选择部署模式', trigger: 'blur' }
          ],
          group_id: [
            { required: true, message: '请选择集群', trigger: 'blur' }
          ],
          deploy_user: [
            { required: true, message: '请输入部署用户', trigger: 'blur' }
          ],
          deploy_path: [
            { required: true, message: '请输入部署路径', trigger: 'blur' }
          ]
        },
        projectId: 0
      }
    },
    created() {
      this.fetchServerGroup()
      this.fetchProject()
    },
    methods: {
      onSubmit(form) {
        this.$refs[form].validate((valid) => {
          if (valid) {
            update(this.projectId, this.form).then(response => {
              this.$message({
                message: response.message,
                type: 'success'
              })

              this.$router.push({ name: 'project.index' })
            })
          } else {
            return false
          }
        })
      },
      fetchServerGroup() {
        getList({
          'brief': true
        }).then(response => {
          this.serverGroupList = response.data
        })
      },
      fetchProject() {
        let projectId = this.$route.params.id
        this.projectId = projectId

        show(projectId).then(response => {
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

