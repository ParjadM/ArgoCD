hs = {}
if obj.status ~= nil then
  if obj.status.conditions ~= nil then
    for _, condition in ipairs(obj.status.conditions) do
      if condition.type == "InvalidSpec" then
        hs.status = "Degraded"
        hs.message = condition.message
        return hs
      end
    end
  end
  if obj.status.currentPodHash ~= nil then
    if obj.spec.replicas ~= nil and obj.status.updatedReplicas < obj.spec.replicas then
      hs.status = "Progressing"
	  hs.message = "Waiting for roll out to finish: More replicas need to be updated"
      return hs
    end

    local paused = false
    if obj.status.verifyingPreview ~= nil then
      paused = obj.status.verifyingPreview
    elseif obj.spec.paused ~= nil then
      paused = obj.spec.paused
    end

    if paused then
      hs.status = "Suspended"
      if obj.status.blueGreen.previewSelector ~= nil and obj.status.blueGreen.previewSelector == obj.status.currentPodHash then
        hs.message = "The preview Service is serving traffic to the current pod spec"
      elseif obj.status.blueGreen.activeSelector ~= nil and obj.status.blueGreen.activeSelector == obj.status.currentPodHash then
        hs.message = "The active Service is serving traffic to the current pod spec"
      end
      return hs
    end

    if obj.status.replicas > obj.status.updatedReplicas then
      hs.status = "Progressing"
      hs.message = "Waiting for roll out to finish: old replicas are pending termination"
      return hs
    end
    if obj.status.availableReplicas < obj.status.updatedReplicas then
      hs.status = "Progressing"
      hs.message = "Waiting for roll out to finish: updated replicas are still becoming available"
      return hs
    end

    if obj.status.blueGreen.activeSelector ~= nil and obj.status.blueGreen.activeSelector == obj.status.currentPodHash then
      hs.status = "Healthy"
      hs.message = "The active Service is serving traffic to the current pod spec"
      return hs
    end
    hs.status = "Progressing"
    hs.message = "The current pod spec is not receiving traffic from the active service"
    return hs
  end
end
hs.status = "Progressing"
hs.message = "Waiting for rollout to finish: status has not been reconciled."
return hs