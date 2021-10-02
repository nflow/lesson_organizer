<template>
  <board :boardId="boardId" v-if="boardId" />
</template>
<script lang="ts">
import { postBoard } from "@/api/board";
import { BoardDto, CreateBoardDto } from "@/types/board";
import { defineComponent, ref } from "vue";
import { useQueryClient } from "vue-query";
import { useRoute, useRouter } from "vue-router";
import Board from "./Board.vue";

export default defineComponent({
  name: "Landing",
  components: {
    Board,
  },
  setup() {
    const router = useRouter();
    const route = useRoute();

    const createBoard = postBoard();
    const queryClient = useQueryClient();

    let boardId = ref(route.params.boardId);

    if (!route.params.boardId) {
      createBoard.mutate(
        {
          name: "New Board",
          goals: [],
          contents: [],
          phases: [],
        } as CreateBoardDto,
        {
          onSuccess: (data: BoardDto) => {
            queryClient.setQueryData(["board"], data);
            boardId.value = data.id;
            router.replace({
              name: "Landing",
              params: {
                boardId: data ? data.id : "",
              },
            });
          },
        }
      );
    }

    return { boardId };
  },
});
</script>
<style></style>
