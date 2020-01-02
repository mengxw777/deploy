<template>
  <div class="app-container">
    <el-row>
      <el-button type="primary" size="mini" style="float: right; margin-bottom: 10px" @click="newDeploy">
        新建申请单
      </el-button>
    </el-row>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="ID">
        <template slot-scope="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>
      <el-table-column label="名称">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="项目" align="center">
        <template slot-scope="scope">
          {{ scope.row.project.name }}
        </template>
      </el-table-column>
      <el-table-column label="申请人" align="center">
        <template slot-scope="scope">
          {{ scope.row.user.account }}
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center" width="170px">
        <template slot-scope="scope">
          {{ parseTime(scope.row.created_at) }}
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <el-row>
            <el-button type="primary" size="mini">
              <router-link :to="{ name: 'deploy.show', params: { id: scope.row.id }}">查看</router-link>
            </el-button>
          </el-row>
          <br/>
          <el-row>
            <el-button type="danger" size="mini" @click="onDelete(scope.row.id)"
                       v-if="scope.row.build_status !== 4 && scope.row.deploy_status !== 2">
              废弃
            </el-button>
          </el-row>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination style="float: right; margin-top: 10px"
                   background
                   layout="prev, pager, next"
                   :total=total
                   :page-sizes=perPage
    >
    </el-pagination>

    <el-dialog title="新申请单" :visible.sync="dialogFormVisible">
      <el-form :model="form" label-position="right" label-width="50px">
        <el-form-item label="名称">
          <el-input v-model="form.name" autocomplete="off"/>
        </el-form-item>
        <el-form-item label="项目">
          <el-select v-model="form.project_id" filterable placeholder="请选择项目">
            <el-option
              v-for="item in projectList"
              :key="item.id"
              :label="item.name"
              :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="onSubmit()">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  import { getList, store, destroy } from '@/api/deploy'
  import { parseTime } from '@/utils/index'
  import { getList as getProjectList } from '@/api/project'

  export default {
    filters: {
      statusFilter(status) {
        const statusMap = {
          published: 'success',
          draft: 'gray',
          deleted: 'danger'
        }
        return statusMap[status]
      }
    },
    data() {
      return {
        list: null,
        listLoading: true,
        perPage: [],
        total: 0,
        visible: true,
        dialogFormVisible: false,
        form: {
          name: '',
          project_id: ''
        },
        projectList: []
      }
    },
    created() {
      this.fetchData()
    },
    methods: {
      fetchData() {
        this.listLoading = true
        getList().then(response => {
          this.list = response.data
          this.total = response.pagination.total
          this.perPage = [response.pagination.per_page]
          this.listLoading = false
        })
      },
      parseTime(data) {
        return parseTime(data)
      },
      onDelete(id) {
        this.$confirm('确认作废此申请单?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {

          destroy(id).then(response => {
            this.$message({
              message: response.message,
              type: 'success'
            })

            this.fetchData()
          })

        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
      },
      newDeploy() {
        this.fetchProject()
        this.dialogFormVisible = this.dialogFormVisible !== true
      },
      fetchProject() {
        getProjectList({
          "brief": true
        }).then(response => {
          this.projectList = response.data
        })
      },
      onSubmit() {
        store(this.form).then(response => {
          this.fetchData()
          this.dialogFormVisible = this.dialogFormVisible !== true
        })
      }
    }
  }
</script>
