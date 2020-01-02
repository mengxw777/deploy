<template>
  <div class="app-container">
    <el-row>
      <el-button type="primary" size="mini" style="float: right; margin-bottom: 10px">
        <router-link :to="{ name: 'server.store',}">添加服务器</router-link>
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
      <el-table-column label="IP地址" align="center">
        <template slot-scope="scope">
          {{ scope.row.ip }}
        </template>
      </el-table-column>
      <el-table-column label="端口" align="center">
        <template slot-scope="scope">
          {{ scope.row.port }}
        </template>
      </el-table-column>
      <el-table-column label="集群" align="center">
        <template slot-scope="scope">
          <router-link :to="{ name: 'serverGroup.index', params: {id: scope.row.group.id}}">
            {{ scope.row.group.name }}
          </router-link>
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
              <router-link :to="{ name: 'server.update', params: { id: scope.row.id }}">编辑</router-link>
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
  import { getList, destroy } from '@/api/server'
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
        this.$confirm('确认删除此服务器?', '提示', {
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
