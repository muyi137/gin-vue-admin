<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="基本工资:" prop="basis">
          <el-input-number v-model="formData.basis" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="身份证号:" prop="card_number">
          <el-input v-model="formData.card_number" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="岗位津贴:" prop="jobs">
          <el-input-number v-model="formData.jobs" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="学历:" prop="eduLevel">
          <el-input-number v-model="formData.eduLevel" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="专业职称:" prop="proTitles">
          <el-input-number v-model="formData.proTitles" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="工作年限:" prop="workYears">
          <el-input-number v-model="formData.workYears" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="统筹:" prop="planWhole">
          <el-input-number v-model="formData.planWhole" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="自定义:" prop="custom1">
          <el-input-number v-model="formData.custom1" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="自定义2:" prop="custom2">
          <el-input-number v-model="formData.custom2" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="自定义3:" prop="custom3">
          <el-input-number v-model="formData.custom3" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="自定义4:" prop="custom4">
          <el-input-number v-model="formData.custom4" :precision="2" :clearable="true"></el-input-number>
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
  name: 'OaSalary'
}
</script>

<script setup>
import {
  createOaSalary,
  updateOaSalary,
  findOaSalary
} from '@/api/oaSalary'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            basis: 0,
            card_number: '',
            jobs: 0,
            eduLevel: 0,
            proTitles: 0,
            workYears: 0,
            planWhole: 0,
            custom1: 0,
            custom2: 0,
            custom3: 0,
            custom4: 0,
        })
// 验证规则
const rule = reactive({
               basis : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               card_number : [{
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
      const res = await findOaSalary({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reoaSalary
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
               res = await createOaSalary(formData.value)
               break
             case 'update':
               res = await updateOaSalary(formData.value)
               break
             default:
               res = await createOaSalary(formData.value)
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
