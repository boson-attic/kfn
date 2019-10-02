FROM fedora:latest

# Don't include container-selinux and remove
# directories used by dnf that are just taking
# up space. Adjust storage.conf to enable Fuse storage.
RUN yum -y install buildah fuse-overlayfs --exclude container-selinux; rm -rf /var/cache /var/log/dnf* /var/log/yum.*; \
    sed -i -e 's|^#mount_program|mount_program|g' -e '/additionalimage.*/a "/var/lib/shared",' /etc/containers/storage.conf; \
    mkdir -p /var/lib/shared/overlay-images /var/lib/shared/overlay-layers; touch /var/lib/shared/overlay-images/images.lock; touch /var/lib/shared/overlay-layers/layers.lock

COPY kfn /usr/bin

# Set up environment variables to note that this is
# not starting with usernamespace and default to
# isolate the filesystem with chroot.
ENV BUILDAH_ISOLATION=chroot
ENV KFN_IN_CLUSTER=true

ENTRYPOINT ["kfn"]
