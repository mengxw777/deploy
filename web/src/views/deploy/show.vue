<template>
  <div class="app-container">
    <el-card class="box-card">
      <div slot="header">
        <span>申请单 / {{ data.name }}</span>
      </div>
      <el-row>
        <el-col :span="12">
          <div class="grid-content bg-purple">项目：{{ data.project.name }}</div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple">上线模式：{{ data.project.deploy_mode === 1 ? '分支' : '标签' }}</div>
        </el-col>
      </el-row>

      <br/>
      <br/>
      <el-row>
        <el-col :span="12">
          <div class="grid-content bg-purple">提交人：{{ data.user.account }}</div>
        </el-col>
      </el-row>

      <br/>
      <br/>
      <el-row>
        <el-col :span="12">
          <div class="grid-content bg-purple">构建状态：
            <span v-if="data.build_status === 0">
              未构建
              <el-button type="primary" size="mini" @click="build">构建</el-button>
            </span>
            <span v-if="data.build_status === 1">构建中</span>
            <span v-if="data.build_status === 2">
              构建成功
             <el-button v-show="data.deploy_status !== 2" type="primary" size="mini" @click="build">重新构建</el-button>
            </span>
            <span v-if="data.build_status === 3">
              构建失败
              <el-button v-show="data.deploy_status !== 2" type="primary" size="mini" @click="build">重新构建</el-button>
            </span>
            <span v-if="data.build_status === 4">已废弃</span>

          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple">压缩包位置：{{ data.build_path }}</div>
        </el-col>
      </el-row>

      <br/>
      <br/>
      <el-row>
        <el-col :span="12">
          <div class="grid-content bg-purple">构建日志：
            <el-button type="primary" size="mini" @click="openBuildLog">查看</el-button>
          </div>
        </el-col>
        <el-col :span="12">

        </el-col>
      </el-row>

      <br/>
      <br/>
      <el-row>
        <el-col :span="12">
          <div class="grid-content bg-purple">部署状态：
            <span v-if="data.deploy_status === 0">
              未部署
              <el-button type="primary" size="mini" @click="deploy" v-if="data.build_status !== 4">部署</el-button>
            </span>
            <span v-if="data.deploy_status === 1">部署中</span>
            <span v-if="data.deploy_status === 2">部署成功</span>
            <span v-if="data.deploy_status === 3">
              部署失败
              <el-button type="primary" size="mini" @click="deploy">重新部署</el-button>
            </span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple">部署日志：
            <el-button type="primary" size="mini" @click="openDeployLog">查看</el-button>
          </div>
        </el-col>
      </el-row>
    </el-card>

    <el-dialog title="构建日志" :visible.sync="buildLogVisible" :append-to-body="true">
      <el-table :data="data.build_log" height="800" stripe>
        <el-table-column property="cmd" label="执行命令"/>
        <el-table-column property="stderr" label="标准输出"/>
        <el-table-column property="stdout" label="执行结果"/>
        <el-table-column property="success" label="执行结果">
          <template slot-scope="scope">
            {{ scope.row.success === true ? '执行成功' : '执行失败' }}
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog title="部署日志" :visible.sync="deployLogVisible" :append-to-body="true">
      <el-table :data="data.deploy_log" height="800" stripe>
        <el-table-column property="cmd" label="执行命令"/>
        <el-table-column property="stderr" label="标准输出"/>
        <el-table-column property="stdout" label="执行输出"/>
        <el-table-column property="success" label="执行结果">
          <template slot-scope="scope">
            {{ scope.row.success === true ? '执行成功' : '执行失败' }}
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
  </div>
</template>

<script>
  import { show, build, deploy } from '@/api/deploy'
  import { parseTime } from '@/utils/index'

  export default {
    data() {
      return {
        deployId: '',
        data: {
          'project': '',
          'user': ''
        },
        buildLogVisible: false,
        deployLogVisible: false,
        loading: null
      }
    },
    created() {
      this.fetchDeploy()
      this.listenWs()
    },
    methods: {
      fetchDeploy() {
        this.deployId = this.$route.params.id

        show(this.deployId).then(response => {
          this.data = response.data
        })
      },
      openDeployLog() {
        this.deployLogVisible = this.deployLogVisible !== true
      },
      openBuildLog() {
        this.buildLogVisible = this.buildLogVisible !== true
      },
      listenWs() {
        let that = this
        this.$options.sockets.onmessage = function(data) {
          let msg = JSON.parse(data.data)
          if (
            msg.action === 'build success' ||
            msg.action === 'build fail' ||
            msg.action === 'deploy success' ||
            msg.action === 'deploy fail'
          ) {
            that.fetchDeploy()
            that.loading.close()
          }
        }
      },
      build() {
        this.loading = this.$loading({
          lock: true,
          text: 'Loading',
          spinner: 'el-icon-loading',
          background: 'rgba(0, 0, 0, 0.7)'
        })

        build(this.deployId)
      },
      deploy() {
        this.loading = this.$loading({
          lock: true,
          text: 'Loading',
          spinner: 'el-icon-loading',
          background: 'rgba(0, 0, 0, 0.7)'
        })

        deploy(this.deployId)
      }
    }
  }
</script>

