<template>
  <div class="app-container">
    <el-row>
      <el-button type="primary" size="mini" style="float: right; margin-bottom: 10px"
                 @click="newGroupDialogVisible = true">
        新建集群
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
      <el-table-column label="创建时间" align="center">
        <template slot-scope="scope">
          {{ parseTime(scope.row.created_at) }}
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <el-dropdown>
            <el-button type="primary" size="mini">
              操作<i class="el-icon-arrow-down el-icon--right"/>
            </el-button>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item @click.native="onUpdate(scope.row.id, scope.row.name)">编辑</el-dropdown-item>
              <el-dropdown-item @click.native="onDelete(scope.row.id)">删除</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination style="float: right; margin-top: 10px"
                   background
                   layout="prev, pager, next"
                   :total=total
                   :page-size=pageSize
                   @current-change="handleCurrentChange"
    >
    </el-pagination>

    <el-dialog title="编辑" :visible.sync="dialogFormVisible">
      <el-form>
        <el-form-item label="集群名称">
          <el-input autocomplete="off" v-model="groupName"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="updateSubmit">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog title="新建集群" :visible.sync="newGroupDialogVisible">
      <el-form>
        <el-form-item label="集群名称">
          <el-input autocomplete="off" v-model="groupName"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="newGroupDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="newGroupSubmit">确 定</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
  import { getList, destroy, update, store } from '@/api/serverGroup'
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
        pageSize: 0,
        total: 0,
        visible: true,
        dialogFormVisible: false,
        groupID: 0,
        groupName: '',

        newGroupDialogVisible: false
      }
    },
    created() {
      this.fetchData()
    },
    methods: {
      fetchData(page) {
        this.listLoading = true

        getList({
          page: page
        }).then(response => {
          this.list = response.data
          this.total = response.pagination.total
          this.pageSize = response.pagination.per_page
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
      },
      onUpdate(id, name) {
        this.dialogFormVisible = true
        this.groupID = id
        this.groupName = name
      },
      updateSubmit() {
        update(this.groupID, {
          name: this.groupName
        }).then(response => {
          this.dialogFormVisible = false
          this.groupID = 0
          this.groupName = ''

          this.fetchData()
        })
      },
      newGroupSubmit() {
        store({
          name: this.groupName
        }).then(response => {
          this.newGroupDialogVisible = false
          this.groupName = ''
          this.fetchData()
        })
      },
      handleCurrentChange(page) {
        this.fetchData(page)
      }
    }
  }
</script>
