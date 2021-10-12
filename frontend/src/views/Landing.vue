<template>
  <div class="tw-bg-red-300"></div>
</template>
<script lang="ts">
import { postBoard } from "@/api/board";
import { BoardDto, CreateBoardDto } from "@/types/board";
import { defineComponent, ref } from "vue";
import { useQueryClient } from "vue-query";
import { useRoute, useRouter } from "vue-router";
export default defineComponent({
  name: "Landing",
  setup() {
    const router = useRouter();
    const route = useRoute();

    const createBoard = postBoard();
    const queryClient = useQueryClient();

    // if (!route.params.boardId) {
    //   createBoard.mutate(
    //     {
    //       name: "New Board",
    //       goals: [],
    //       contents: [],
    //       phases: [],
    //     } as CreateBoardDto,
    //     {
    //       onSuccess: (data: BoardDto) => {
    //         queryClient.setQueryData(["board"], data);
    //         boardId.value = data.id;
    //         router.replace({
    //           name: "Landing",
    //           params: {
    //             boardId: data ? data.id : "",
    //           },
    //         });
    //       },
    //     }
    //   );
    // }

    return {};
  },
});
</script>
<style></style>
