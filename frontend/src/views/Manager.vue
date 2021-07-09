<template>
  <div class="flex flex-col bg-gray-200 m-0 h-full">
    <div
      class="
        flex flex-col flex-initial
        bg-gradient-to-br
        from-blue
        to-turquoise
        p-4
        shadow-md
        space-y-4
      "
    >
      <div class="flex flex-row gap-1 flex-auto space-x-2">
        <div
          class="
            flex flex-initial flex-col
            items-center
            p-3
            h-full
            rounded
            text-white
            shadow-xl
            bg-gradient-to-br
            from-red-300
            to-orange
          "
        >
          <span class="flex-auto text-xl text-center font-logo"
            >Lesson Organizer - Manager
          </span>
        </div>
      </div>
    </div>
    <div class="flex flex-col flex-auto p-5">
      <h1
        class="
          text-xl text-white
          font-bold
          p-2
          rounded
          shadow
          bg-gradient-to-r
          from-blue
          to-turquoise
        "
      >
        Methoden Verwaltung
      </h1>
      <div class="mt-2 mb-2 p-2 bg-gray-300 rounded">
        <el-form>
          <el-button>Test</el-button>
          <el-input type="textarea" />
        </el-form>
      </div>
    </div>
  </div>
</template>
<script>
import { defineComponent } from "@vue/runtime-core";

export default defineComponent({
  setup() {
    return {};
  },
});
</script>
