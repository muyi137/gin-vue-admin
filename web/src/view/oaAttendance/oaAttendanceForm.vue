<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="身份证号:" prop="cardNumber">
          <el-input v-model="formData.cardNumber" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="月份:" prop="month">
          <el-input v-model.number="formData.month" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="应出勤天数:" prop="need">
          <el-input-number v-model="formData.need" :precision="2" :clearable="false"></el-input-number>
        </el-form-item>
        <el-form-item label="请假天数:" prop="leave">
          <el-input-number v-model="formData.leave" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="旷工:" prop="absent">
          <el-input-number v-model="formData.absent" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="迟到早退天数:" prop="cdzt">
          <el-input-number v-model="formData.cdzt" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="状态:" prop="status">
        <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['已审核','待审核']" :key="item" :label="item" :value="item" />
        </el-select>
        </el-form-item>
        <el-form-item label="实际签到天数:" prop="signed">
          <el-input v-model="formData.signed" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'OaAttendance'
}
</script>

<script setup>
import {
  createOaAttendance,
  updateOaAttendance,
  findOaAttendance
} from '@/api/oaAttendance'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            cardNumber: '',
            month: 0,
            need: 0,
            leave: 0,
            absent: 0,
            cdzt: 0,
            signed: '',
        })
// 验证规则
const rule = reactive({
               cardNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               month : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               need : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               signed : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findOaAttendance({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reoaAttendance
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createOaAttendance(formData.value)
               break
             case 'update':
               res = await updateOaAttendance(formData.value)
               break
             default:
               res = await createOaAttendance(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
