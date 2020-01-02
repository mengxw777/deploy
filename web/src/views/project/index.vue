<template>
  <div class="app-container">
    <el-row>
      <el-button type="primary" size="mini" style="float: right; margin-bottom: 10px">
        <router-link :to="{ name: 'project.store',}">新建项目</router-link>
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
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="Git 地址">
              <span>{{ props.row.git_addr }}</span>
            </el-form-item>
            <el-form-item label="构建命令">
              <span>{{ props.row.build_cmd }}</span>
            </el-form-item>
            <el-form-item label="部署用户">
              <span>{{ props.row.deploy_user }}</span>
            </el-form-item>
            <el-form-item label="部署目录">
              <span>{{ props.row.deploy_path }}</span>
            </el-form-item>
            <el-form-item label="部署前命令">
              <span>{{ props.row.deploy_before_cmd }}</span>
            </el-form-item>
            <el-form-item label="部署后命令">
              <span>{{ props.row.deploy_after_cmd }}</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column align="center" label="ID">
        <template slot-scope="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>
      <el-table-column label="项目名称">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="部署模式" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.deploy_mode === 1 ? '分支' : '标签' }}</span>
        </template>
      </el-table-column>
      <el-table-column label="集群" align="center">
        <template slot-scope="scope">
          {{ scope.row.server_group.name }}
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center">
        <template slot-scope="scope">
          {{ parseTime(scope.row.created_at) }}
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <el-row>
            <el-button type="primary" size="mini">
              <router-link :to="{ name: 'project.update', params: { id: scope.row.id }}">编辑</router-link>
            </el-button>
          </el-row>
          <br/>
          <el-row>
            <el-button type="danger" size="mini" @click="onDelete(scope.row.id)">删除</el-button>
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
  </div>
</template>

<script>
  import { getList, destroy } from '@/api/project'
  import { parseTime } from '@/utils/index'

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
        visible: true
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
        this.$confirm('确认删除此项目?', '提示', {
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
      }
    }
  }
</script>

<style>
  .demo-table-expand {
    font-size: 0;
  }

  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }

  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
</style>
